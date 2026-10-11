package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Mega Project Success! Frontend is alive and running bava!")
    })
    fmt.Println("Server is running on port 80...")
    http.ListenAndServe(":80", nil) // Intha line thaan program-a nillama run panna vekkum
}
