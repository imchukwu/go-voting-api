package main

import (
    "go-voting-api/config"
    "go-voting-api/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    config.ConnectDatabase()
    routes.RegisterRoutes(r)

    r.Run() // default is :8080
}
