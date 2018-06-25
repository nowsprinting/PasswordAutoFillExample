package backend

import "net/http"

func init() {
	// Login form on root directory
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static_files/login.html")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static_files/login_complete.html")
	})

	// Login form on sub directory
	http.HandleFunc("/sub/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static_files/login.html")
	})
	http.HandleFunc("/sub/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static_files/login_complete.html")
	})

	// for iOS apps
	http.HandleFunc("/.well-known/apple-app-site-association", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // `curl --verbose`で有効なのを確認済
		http.ServeFile(w, r, "static_files/apple-app-site-association.json")
	})
}
