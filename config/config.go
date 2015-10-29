package config

import(
    "github.com/rs/cors"
)

const (
    HOST = "0.0.0.0"
    PORT = ":3000"
)

var CORS = cors.New(cors.Options{
    AllowedHeaders: []string{"Access-Token", "X-Requested-With", "Content-Type", "Accept"},
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
})
