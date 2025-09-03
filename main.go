package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/leventeberry/goapi/initializers"
    "github.com/leventeberry/goapi/routes"
)

func main() {
    // Initialize environment variables, database connection, and run migrations
    initializers.Init()

    // Create a Gin router with default middleware (logger and recovery)
    router := gin.Default()

    // Register all routes
    routes.SetupRoutes(router, initializers.DB)

    // Start server on specified PORT or default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    if err := router.Run(":" + port); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}