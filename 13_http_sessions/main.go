// ============================================================
// LESSON 13: HTTP Server with Sessions & Cookies
// ============================================================
// Build web servers with Go's built-in net/http package
// Handle sessions, cookies, JSON, and routing
//
// TO RUN: go run main.go
// Then open http://localhost:8080 in your browser
// ============================================================

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

// ==========================================
// IN-MEMORY SESSION STORE
// ==========================================

type Session struct {
	Username  string
	LoginTime time.Time
	Data      map[string]string
}

var (
	sessions = make(map[string]*Session)
	mu       sync.RWMutex // For thread-safe access
)

// Generate simple session ID (use UUID in production!)
func generateSessionID() string {
	return fmt.Sprintf("sess_%d", time.Now().UnixNano())
}

// Get session from cookie
func getSession(r *http.Request) *Session {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil
	}
	mu.RLock()
	defer mu.RUnlock()
	return sessions[cookie.Value]
}

// Create new session
func createSession(w http.ResponseWriter, username string) *Session {
	sessionID := generateSessionID()
	session := &Session{
		Username:  username,
		LoginTime: time.Now(),
		Data:      make(map[string]string),
	}

	mu.Lock()
	sessions[sessionID] = session
	mu.Unlock()

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   3600, // 1 hour
		HttpOnly: true,
	})

	return session
}

// Delete session
func deleteSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		mu.Lock()
		delete(sessions, cookie.Value)
		mu.Unlock()
	}

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

// ==========================================
// HTTP HANDLERS
// ==========================================

// Home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	session := getSession(r)

	html := `<!DOCTYPE html>
<html>
<head><title>Go HTTP Session Demo</title>
<style>
    body { font-family: Arial; max-width: 600px; margin: 50px auto; padding: 20px; }
    .card { background: #f5f5f5; padding: 20px; border-radius: 8px; margin: 20px 0; }
    a { color: #007bff; }
    input, button { padding: 10px; margin: 5px 0; }
    button { background: #007bff; color: white; border: none; cursor: pointer; }
</style>
</head>
<body>
    <h1>🚀 Go HTTP Session Demo</h1>`

	if session != nil {
		html += fmt.Sprintf(`
    <div class="card">
        <h2>Welcome, %s!</h2>
        <p>Logged in at: %s</p>
        <p><a href="/dashboard">Go to Dashboard</a></p>
        <p><a href="/logout">Logout</a></p>
    </div>`, session.Username, session.LoginTime.Format("3:04 PM"))
	} else {
		html += `
    <div class="card">
        <h2>Login</h2>
        <form action="/login" method="POST">
            <input type="text" name="username" placeholder="Username" required><br>
            <button type="submit">Login</button>
        </form>
    </div>`
	}

	html += `
    <h3>API Endpoints:</h3>
    <ul>
        <li><a href="/api/time">/api/time</a> - Get current time (JSON)</li>
        <li><a href="/api/users">/api/users</a> - Get users (JSON)</li>
    </ul>
</body></html>`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

// Login handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Username required", http.StatusBadRequest)
		return
	}

	createSession(w, username)
	log.Printf("User '%s' logged in", username)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout handler
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session := getSession(r)
	if session != nil {
		log.Printf("User '%s' logged out", session.Username)
	}
	deleteSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Dashboard (protected route)
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	session := getSession(r)
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head><title>Dashboard</title></head>
<body>
    <h1>Dashboard</h1>
    <p>Hello, %s! This is a protected page.</p>
    <p>Session started: %s</p>
    <p><a href="/">Home</a> | <a href="/logout">Logout</a></p>
</body></html>`, session.Username, session.LoginTime.Format(time.RFC1123))

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

// ==========================================
// JSON API ENDPOINTS
// ==========================================

// API: Get current time
func apiTimeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"time":      time.Now().Format(time.RFC3339),
		"timestamp": time.Now().Unix(),
		"timezone":  time.Now().Location().String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// API: Get users (demo)
func apiUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []map[string]interface{}{
		{"id": 1, "name": "Alice", "email": "alice@example.com"},
		{"id": 2, "name": "Bob", "email": "bob@example.com"},
		{"id": 3, "name": "Charlie", "email": "charlie@example.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// ==========================================
// MIDDLEWARE
// ==========================================

// Logging middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("→ %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("← %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	}
}

// ==========================================
// MAIN
// ==========================================

func main() {
	// Routes
	http.HandleFunc("/", loggingMiddleware(homeHandler))
	http.HandleFunc("/login", loggingMiddleware(loginHandler))
	http.HandleFunc("/logout", loggingMiddleware(logoutHandler))
	http.HandleFunc("/dashboard", loggingMiddleware(dashboardHandler))
	http.HandleFunc("/api/time", loggingMiddleware(apiTimeHandler))
	http.HandleFunc("/api/users", loggingMiddleware(apiUsersHandler))

	// Start server
	port := ":8080"
	fmt.Println("===========================================")
	fmt.Println("🚀 Go HTTP Server with Sessions")
	fmt.Println("===========================================")
	fmt.Printf("Server running at http://localhost%s\n", port)
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println("===========================================")

	log.Fatal(http.ListenAndServe(port, nil))
}

// Unused but useful: template example
var _ = template.Must(template.New("example").Parse(`{{.}}`))
