package main

import (
    "app/lib"
    "app/config"

    "log"
    "net/http"
)

func main() {

    router := lib.NewRouter()

    log.Printf("Server Listening on localhost%s", config.PORT)
    log.Fatal(http.ListenAndServe(config.PORT, router))

}