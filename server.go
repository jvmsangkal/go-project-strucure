package main

import (
    "app/lib"
    "app/config"

    "log"
    "net/http"
)

func main() {

    router := lib.NewRouter()

    handler := config.CORS.Handler(router)

    log.Printf("Server Listening on %s", config.HOST + config.PORT)
    log.Fatal(http.ListenAndServe(config.HOST + config.PORT, handler))

}