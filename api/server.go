package api

import (
	"log"

	"github.com/Jojojojodr/taskflow/api/routers"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func StartServer(port string) {
	log.Println("Starting server...")

	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server is running on port %s\n", port)
	
	svr := gin.Default()
	err := svr.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	routers.SetupRoutes(svr)

	log.Printf("Server is running on http://localhost:%s\n", port)
	svr.Use(cors.Default())
	svr.Run(":" + port)
}