package routes

import (
	"ecommerce-app/api/controllers"
	"ecommerce-app/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)

	admin := r.Group("/admin")
	admin.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	admin.POST("/products", controllers.CreateProduct)
	admin.PUT("/products/:id", controllers.UpdateProduct)
	admin.DELETE("/products/:id", controllers.DeleteProduct)
}
