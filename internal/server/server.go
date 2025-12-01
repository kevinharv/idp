package server

import (
	"log"
	"net/http"
	"github.com/kevinharv/idp/internal/handlers/auth"
)

func Serve() {
	mux := http.NewServeMux()
	oauthMux := http.NewServeMux()

	loadRoutes(oauthMux)
	// Ensure oauth subpaths are routed to oauthMux (longest pattern wins)
	mux.Handle("/oauth/", oauthMux)
	mux.HandleFunc("/login", auth.GetLoginPage)

	// Serve files from web/static at the site root.
	// Example: GET /css/app.css -> web/static/css/app.css
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/", fs)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
