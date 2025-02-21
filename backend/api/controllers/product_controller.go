package controllers

import (
	"ecommerce-app/database"
	"ecommerce-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// Create a new product (Admin Only)
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// Update product details (Admin Only)
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// Delete a product (Admin Only)
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Product{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
