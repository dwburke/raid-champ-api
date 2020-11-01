package champ

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	//r.HandleFunc("/champ", ListChamp).Methods("GET")
	//r.HandleFunc("/champ/{id}", GetChamp).Methods("GET")
	r.HandleFunc("/champ", CreateChamp).Methods("POST")
	//r.HandleFunc("/champ/{id}", UpdateChamp).Methods("PATCH")
	//r.HandleFunc("/champ/{id}", DeleteChamp).Methods("DELETE")
}
