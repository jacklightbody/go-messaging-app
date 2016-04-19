package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
	// pretty much just a router
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", messagePost).Methods("POST")
    router.HandleFunc("/{from}/{to}/{fromTimestamp}", messagesIndex).Methods("GET")
    router.HandleFunc("/{from}/{to}/", timelessMessagesIndex).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}
