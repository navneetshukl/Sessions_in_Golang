package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Create a new session store
var store = sessions.NewCookieStore([]byte("secret-key"))

// Middleware to check if user is authenticated
// Middleware to check if user is authenticated
func isAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Check if user is authenticated
		if session.Values["authenticated"] != true {
			fmt.Println("No Session")
			//http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// Login handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// Set user as authenticated
	fmt.Println("Login Route")
	session.Values["authenticated"] = true
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}

// Logout handler
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// Revoke user authentication
	session.Values["authenticated"] = false
	session.Save(r, w)

	//http.Redirect(w, r, "/login", http.StatusFound)
}

// Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Route")
	fmt.Fprintln(w, "Welcome to the home page!")
}

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/login", loginHandler).Methods("GET")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")
	r.HandleFunc("/", isAuthenticated(homeHandler)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
