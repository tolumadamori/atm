package routes

import (
	"github.com/gorilla/mux"
	"github.com/tolumadamori/atm/pkg/controllers"
)

func Router(r *mux.Router) {
	r.HandleFunc("/registeruser/", controllers.CreateUser).Methods("Post")   //creates a user
	r.HandleFunc("/openAccount/", controllers.CreateAccount).Methods("Post") //creates an account
	r.HandleFunc("/deposit/", controllers.Deposit).Methods("Post")
	r.HandleFunc("/withdraw/", controllers.Withdraw).Methods("Post")
	r.HandleFunc("/balance/", controllers.Balance).Methods("Get")
	// r.HandleFunc("/closeaccount/", controllers.CloseAccount).Methods("Post")

}
