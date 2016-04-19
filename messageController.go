package main

import (
    "strings"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func messagePost(w http.ResponseWriter, r *http.Request) {
    from := r.FormValue("from")
    to := r.FormValue("to")
    body := r.FormValue("message")
    if (len(strings.TrimSpace(from)) == 0 ||
        len(strings.TrimSpace(to)) == 0 ||
        len(strings.TrimSpace(body)) == 0){
        // invalid input
        json.NewEncoder(w).Encode("Error: Bad Input")
    }else{
        // valid input
        NewMessage(from, to, body)
    }
}

func messagesIndex(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    from := vars["from"]
    to := vars["to"]
    fromTimestamp, err := strconv.Atoi(vars["fromTimestamp"])// convert timestamp to int
    if err != nil { // if the conversion failed just return all messages
        fromTimestamp = 0
    }
    // call our model as normal
    messagesData := GetMessagesAfter(from, to, fromTimestamp)
    // return the data
    json.NewEncoder(w).Encode(messagesData)
}
func timelessMessagesIndex(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    from := vars["from"]
    to := vars["to"]
    // fromTimestamp of 0 will return all messages
    messagesData := GetMessagesAfter(from, to, 0)
    // return the data
    json.NewEncoder(w).Encode(messagesData)
}
