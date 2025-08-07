package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/orders", GetOrders).Methods("GET")
    log.Println("Order service running on port 8085")
    log.Fatal(http.ListenAndServe(":8085", r))
}