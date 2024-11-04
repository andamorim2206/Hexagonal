package main

import (
	"github.com/andamorim2206/hexagonal_1/adapter/input/controllers/routes"
	"github.com/andamorim2206/hexagonal_1/configuration/logger"
	"github.com/gin-gonic/gin"
)

func main() {
    logger.Info("About to start application")
    router := gin.Default()

    routes.InitRoutes(router)

    if err := router.Run(":8080"); err != nil {
        logger.Error("Error trying to init application on port 8080", err)
        return
    }
}
