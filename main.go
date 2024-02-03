package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Set to release mode
	gin.SetMode(gin.ReleaseMode)

	// Set CORS
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"POST", "GET"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Set Router
	processes := router.Group("/processes")
	processes.GET("", GetProcessesController)
	processes.GET("/:id", GetProcessByPIDController)
	processes.POST("/:id/kill", KillProcessByPIDController)
	processes.POST("/kill", KillAllProcessesController)
	router.Run("localhost:4000")

}
