package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/login", Login).Methods("POST")
    r.HandleFunc("/register", Register).Methods("POST")
    log.Println("Auth service running on port 8083")
    log.Fatal(http.ListenAndServe(":8083", r))
}