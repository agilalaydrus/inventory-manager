package main

import (
	"github.com/gin-gonic/gin"
	"inventory-management/src/config"
	"inventory-management/src/routes"
)

func main() {
	config.Connect() // Menghubungkan ke database
	r := gin.Default()
	routes.SetupRoutes(r) // Mengatur rute
	r.Run(":8080")        // Menjalankan server di port 8080
}
