package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"log/slog"
	"net/http"
	"time"
)

func HandleGETOAuthAuthorize(w http.ResponseWriter, r *http.Request) {
	
	// Generate CSRF token and set verification token
	csrfToken, err := generateCSRFToken(32)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO - set to true when HTTPS
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(5 * time.Minute),
	})

	// Load login page template
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		CSRFToken string
	}{CSRFToken: csrfToken}

	// Render HTML template with CSRF token
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "render error", http.StatusInternalServerError)
		return
	}
}

func HandlePOSTOAuthAuthorize(w http.ResponseWriter, r *http.Request) {
	// Read form details - confirm cookie token matches form token
	username := r.FormValue("username")
	password := r.FormValue("password")
	csrfForm := r.FormValue("csrf_token")
	csrfCookie, err := r.Cookie("csrf_token")

	// Bad request if CSRF cookie not provided
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Bad request if CSRF tokens do not match
	if csrfForm != csrfCookie.Value {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	// Auth the user against AD/store
	mockUsername := "test"
	mockPassword := "test"

	slog.Debug("Authorizing " + username)

	if username == mockUsername && password == mockPassword {
		slog.Debug("Successfully authorized " + username)
		http.Redirect(w, r, "/lol", http.StatusFound)
	} else {
		slog.Debug("Failed to authorize " + username)
		w.WriteHeader(http.StatusBadRequest)
	}
}

/* ========= HELPER FUNCTIONS ========= */

func generateCSRFToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
