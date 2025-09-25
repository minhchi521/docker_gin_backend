package main

import (
	"docker-gin-crud/routers"
	"os"
)

func main() {
	r := routers.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
