package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("Hello World ..!!")
	hmux := mux.NewRouter()
	readJSON()
	hmux.HandleFunc("/Employee", getEmployees).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8000", hmux))
}
