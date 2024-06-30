package routes

import (
	"belajar-golang/cmd/ProductService/controllers"

	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func RunRoute() {
	setupRoute()

	Router.Run(":8080")
}

func setupRoute() {
	Router.GET("/products", controllers.GetProducts)
	Router.GET("/product/:id", controllers.GetProductById)
	Router.POST("/product", controllers.CreateProduct)
	Router.PUT("/product/:id", controllers.UpdateProduct)
	Router.DELETE("product/:id", controllers.DeleteProduct)
}
