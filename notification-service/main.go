package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Notification sent to user.")
}

func main() {
    http.HandleFunc("/notify", handler)
    http.ListenAndServe(":8084", nil)
}