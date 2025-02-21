package routes

import (
	"ecommerce-app/api/controllers"
	"ecommerce-app/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.GET("/orders", middlewares.AuthMiddleware(), middlewares.AdminMiddleware(), controllers.GetOrders)
	r.POST("/orders", middlewares.AuthMiddleware(), controllers.CreateOrder)
	r.PUT("/orders/:id", middlewares.AuthMiddleware(), middlewares.AdminMiddleware(), controllers.UpdateOrderStatus)
}
