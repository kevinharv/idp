package server

import (
	"log"
	"net/http"

	"github.com/kevinharv/idp/internal/handlers/oauth"
)

func Serve() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /oauth/authorize", oauth.HandleGETOAuthAuthorize)
	mux.HandleFunc("POST /oauth/authorize", oauth.HandlePOSTOAuthAuthorize)

	// Serve files from web/static at the site root.
	// Example: GET /css/app.css -> web/static/css/app.css
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/", fs)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
