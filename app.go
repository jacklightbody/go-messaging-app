package main

import (
    "fmt"
    "time"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", index)
    router.HandleFunc("/", messagePost).Methods("POST")
    router.HandleFunc("/{uNameA}/{uNameB}/{fromTimestamp}", messagesIndex).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}
