package main

import (
    "rmock/config"
    "rmock/services/user"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Create a new Gin router
    router := gin.Default()

    // Register routes
    user.RegisterRoutes(router)

    // Start the server
    router.Run(":" + cfg.ServerPort)
}
