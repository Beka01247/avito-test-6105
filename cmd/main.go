package main

import (
	"log"
	"os"
	"zadanie-6105/internal/config"
	"zadanie-6105/internal/db"
	"zadanie-6105/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
		db.Connect()

    router := gin.Default()
	  routes.SetupRoutes(router)

    port := os.Getenv("SERVER_ADDRESS")
    if port != "" {
      port = "8080"
    }

	  log.Fatal(router.Run(":"+port))
}
