package server

import (
	"encoding/json"
	"net/http"

	render "mines/app/views"
	"mines/config"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Serve static files from /assets -> public/
	fileServer := http.StripPrefix("/assets/", http.FileServer(http.Dir("public")))
	mux.Handle("/assets/", fileServer)

	// Routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := render.Index(w, map[string]interface{}{
			"AppName": config.GlobalConfig.AppName,
			"Title":   "Modern web apps, the simple way.",
		}, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		err := render.About(w, map[string]interface{}{
			"AppName": config.GlobalConfig.AppName,
			"Title":   "About us",
			"Name":    "Leon",
		}, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/health", s.healthHandler)

	return mux
}

// Health function — now uses stdlib JSON
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s.db.Health()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
