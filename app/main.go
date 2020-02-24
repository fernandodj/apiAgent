package main

import ( 
		"fmt"
		"github.com/gorilla/mux"
		"net/http"
	)


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/save", saveInfo).Methods("POST")
	fmt.Println("Listening on 80");
	http.ListenAndServe(":80", router)
}

