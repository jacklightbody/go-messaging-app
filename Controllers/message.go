package main

import (
    "fmt"
    "time"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the index!")
}
func messagePost(w http.ResponseWriter, r *http.Request) {
    from := r.FormValue("from")
    to := r.FormValue("to")
    body := r.FormValue("message")
    t := time.Now()
    // ...
    insert := Message{From: from, To: to, Body: body, Timestamp: t}
    json.NewEncoder(w).Encode(insert)
}

func messagesIndex(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    uNameA := vars["uNameA"]
    uNameB := vars["uNameB"]
    //fromTimestamp := vars["fromTimestamp"]
    messagesData := Messages{
        Message{From: uNameA, To: uNameB, Body: "Hello"},
        Message{From: uNameB, To: uNameA, Body: "Hi back"},
    }
    json.NewEncoder(w).Encode(messagesData)
}