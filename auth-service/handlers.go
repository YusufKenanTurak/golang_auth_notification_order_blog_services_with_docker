package main

import (
    "encoding/json"
    "net/http"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    json.NewDecoder(r.Body).Decode(&creds)
    if creds.Username == "admin" && creds.Password == "1234" {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Login successful"))
    } else {
        w.WriteHeader(http.StatusUnauthorized)
    }
}

func Register(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User registered"))
}