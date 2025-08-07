package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Post struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Body  string `json:"body"`
}

var posts = []Post{
    {ID: 1, Title: "Hello Go", Body: "Intro to Golang"},
    {ID: 2, Title: "Microservices", Body: "Using Go for microservices"},
}

func main() {
    http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(posts)
    })
    log.Println("Blog service running on :8086")
    log.Fatal(http.ListenAndServe(":8086", nil))
}