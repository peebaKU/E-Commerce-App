package controllers

import (
	"ecommerce-app/database"
	"ecommerce-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all orders (Admin Only)
func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

// Create a new order
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}

// Update order status (Admin Only)
func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&order)
	c.JSON(http.StatusOK, order)
}
