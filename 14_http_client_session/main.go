// ============================================================
// LESSON 14: HTTP Client with Sessions (like Python requests.Session)
// ============================================================
// This is the Go equivalent of Python's requests.Session()
// - Persists cookies across requests
// - Reuses TCP connections
// - Custom headers, timeouts, etc.
//
// TO RUN: go run main.go
// ============================================================

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// ==========================================
// HTTP CLIENT WITH SESSION (Cookie Jar)
// ==========================================

// Session wraps http.Client with helper methods
type Session struct {
	Client  *http.Client
	Headers map[string]string
	BaseURL string
}

// NewSession creates a new session (like requests.Session())
func NewSession() *Session {
	// Create cookie jar to persist cookies
	jar, _ := cookiejar.New(nil)

	return &Session{
		Client: &http.Client{
			Jar:     jar, // This stores cookies automatically!
			Timeout: 30 * time.Second,
		},
		Headers: map[string]string{
			"User-Agent": "Go-Session/1.0",
		},
	}
}

// GET request
func (s *Session) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add session headers
	for k, v := range s.Headers {
		req.Header.Set(k, v)
	}

	return s.Client.Do(req)
}

// POST with form data
func (s *Session) PostForm(url string, data url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range s.Headers {
		req.Header.Set(k, v)
	}

	return s.Client.Do(req)
}

// POST with JSON
func (s *Session) PostJSON(url string, data interface{}) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range s.Headers {
		req.Header.Set(k, v)
	}

	return s.Client.Do(req)
}

// Get cookies for a URL
func (s *Session) GetCookies(rawURL string) []*http.Cookie {
	u, _ := url.Parse(rawURL)
	return s.Client.Jar.Cookies(u)
}

// Set a cookie
func (s *Session) SetCookie(rawURL, name, value string) {
	u, _ := url.Parse(rawURL)
	s.Client.Jar.SetCookies(u, []*http.Cookie{
		{Name: name, Value: value},
	})
}

// Helper: Read response body as string
func ReadBody(resp *http.Response) string {
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

// ==========================================
// EXAMPLE USAGE
// ==========================================

func main() {
	fmt.Println("===========================================")
	fmt.Println("🐍 Go HTTP Session (like Python requests.Session)")
	fmt.Println("===========================================")

	// Create session - equivalent to: session = requests.Session()
	session := NewSession()

	// Set default headers (persist across requests)
	session.Headers["Accept"] = "application/json"
	session.Headers["X-Custom-Header"] = "MyValue"

	fmt.Println("\n=== Example 1: Simple GET Request ===")
	resp, err := session.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status:", resp.Status)
	fmt.Println("Response preview:", ReadBody(resp)[:200], "...")

	fmt.Println("\n=== Example 2: GET with Query Params ===")
	resp, err = session.Get("https://httpbin.org/get?name=Gopher&lang=Go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status:", resp.Status)

	fmt.Println("\n=== Example 3: POST Form Data ===")
	formData := url.Values{
		"username": {"gopher"},
		"password": {"secret123"},
	}
	resp, err = session.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status:", resp.Status)
	fmt.Println("Response preview:", ReadBody(resp)[:300], "...")

	fmt.Println("\n=== Example 4: POST JSON ===")
	jsonData := map[string]interface{}{
		"name":  "Gopher",
		"age":   10,
		"langs": []string{"Go", "Python"},
	}
	resp, err = session.PostJSON("https://httpbin.org/post", jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status:", resp.Status)

	fmt.Println("\n=== Example 5: Cookies Persist Across Requests ===")

	// First request sets a cookie
	session.Get("https://httpbin.org/cookies/set/session_id/abc123")

	// Cookie is automatically sent with next request!
	resp, _ = session.Get("https://httpbin.org/cookies")
	fmt.Println("Cookies in response:", ReadBody(resp))

	// View cookies in our jar
	fmt.Println("Cookies in session jar:")
	for _, cookie := range session.GetCookies("https://httpbin.org") {
		fmt.Printf("  %s = %s\n", cookie.Name, cookie.Value)
	}

	fmt.Println("\n=== Example 6: Manual Cookie ===")
	session.SetCookie("https://httpbin.org", "my_token", "xyz789")
	resp, _ = session.Get("https://httpbin.org/cookies")
	fmt.Println("After adding manual cookie:", ReadBody(resp))

	fmt.Println("\n===========================================")
	fmt.Println("✅ Session demo complete!")
	fmt.Println("===========================================")
}
