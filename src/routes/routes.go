package routes

import (
	"github.com/gin-gonic/gin"
	"inventory-management/src/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/products", controllers.AddProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.POST("/inventory", controllers.CreateInventory)
	r.POST("/inventory/get", controllers.GetInventory)
	r.PUT("/inventory", controllers.UpdateInventory)
	r.DELETE("/inventory", controllers.DeleteInventory)
	r.POST("/inventory/reduce", controllers.ReduceQuantity)
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders/:id", controllers.GetOrderByID)
}
