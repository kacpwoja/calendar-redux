package main

import (
    "fmt"
    "net/http"
    "log"

    "github.com/rs/cors"

    "github.com/kacpwoja/calendar-redux/server/router"
    "github.com/kacpwoja/calendar-redux/server/eventbase"
)

func main() {
    // Alow CORS
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    })

    // Initialize PostgreSQL connection
    db, err := eventbase.Init()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Start on port 4000
    fmt.Println("Starting on port 4000")
    router := router.Router()
    log.Fatal(http.ListenAndServe(":4000", c.Handler(router)))
}