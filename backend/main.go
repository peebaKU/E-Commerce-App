package main

import (
	"ecommerce-app/database"
	"ecommerce-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()
	routes.AuthRoutes(r)
	routes.ProductRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":8080")
}
