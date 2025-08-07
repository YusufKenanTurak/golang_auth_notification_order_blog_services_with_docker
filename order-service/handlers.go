package main

import (
    "encoding/json"
    "net/http"
)

type Order struct {
    ID    string `json:"id"`
    Total string `json:"total"`
}

var orders = []Order{
    {ID: "101", Total: "$300"},
    {ID: "102", Total: "$450"},
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(orders)
}