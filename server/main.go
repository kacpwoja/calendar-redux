package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func hellofunc(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello\n")
}

func main() {
    fmt.Println("Starting on port 4000")
    router := mux.NewRouter()

    router.HandleFunc("/api/", hellofunc).Methods("GET", "OPTIONS")

    log.Fatal(http.ListenAndServe(":4000", router))
}
