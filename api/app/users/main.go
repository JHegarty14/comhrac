package users

import (
	"github.com/gorilla/mux"
)

// Router => alias for the app router so the compiler stops yelling
type Router struct {
	*mux.Router
}

// Creates routes and assigns handlers for users "subrouter"
func (r *Router) addUsersRouterHandlers() {
	r.HandleFunc("/users/auth", LoginHandler).Methods("POST")
	r.HandleFunc("/users/auth/logout", LogoutHandler).Methods("GET")
	r.HandleFunc("/users/signup", RegisterHandler).Methods("POST")
	r.HandleFunc("/users/settings", UserSettingsHandler).Methods("GET")
}
