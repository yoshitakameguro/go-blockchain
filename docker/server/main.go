package main

import (
    "log"
    "net/http"
    "server/db"
    "server/router"
)

func main() {
    db.Init()
    log.Fatal(http.ListenAndServe(":3001", router.Init()))
}
