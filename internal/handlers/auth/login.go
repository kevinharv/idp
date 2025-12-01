package auth

import (
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"net/http"
	"time"
)

// TODO - endpoint for logging in - interactive and programatic

// GET /login - return login page with CSRF token
func GetLoginPage(w http.ResponseWriter, r *http.Request) {
	// generate a secure random token
	token, err := generateCSRFToken(32)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// set CSRF cookie. HttpOnly is fine because we render the token server-side
	// In production, set Secure: true when serving over HTTPS.
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(time.Hour),
	})

	// parse and execute template
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		CSRFToken string
	}{CSRFToken: token}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "render error", http.StatusInternalServerError)
		return
	}
}

// POST /login - check CSRF token, process credentials, continue flow

// generateCSRFToken returns a base64-url encoded random token of n bytes
func generateCSRFToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
