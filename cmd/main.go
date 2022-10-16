package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tolumadamori/atm/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.Router(r)
	http.Handle("/", r)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("failed to start server")
	}

}
