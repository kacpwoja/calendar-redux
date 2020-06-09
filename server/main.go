package main

import (
    "fmt"
    "net/http"
    "log"

    "github.com/kacpwoja/calendar-redux/server/router"
    "github.com/kacpwoja/calendar-redux/server/eventbase"
)

func main() {
    db, err := eventbase.Init()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    fmt.Println("Starting on port 4000")
    router := router.Router()

    log.Fatal(http.ListenAndServe(":4000", router))
}