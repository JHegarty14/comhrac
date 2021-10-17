package main

import "net/http"

// HomepageHandler => Handler for app landing page. Really just checks auth status
func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
