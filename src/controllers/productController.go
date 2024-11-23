package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"inventory-management/src/config"
	"inventory-management/src/models"
	"log"
	"net/http"
)

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := config.DB.Exec("INSERT INTO Products (name, description, price, category) VALUES (?, ?, ?, ?)", product.Name, product.Description, product.Price, product.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})
}

func GetProducts(c *gin.Context) {
	rows, err := config.DB.Query("SELECT * FROM Products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id") // Mengambil ID dari parameter URL

	var product models.Product
	err := config.DB.QueryRow("SELECT * FROM Products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	var requestBody struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	// Bind the JSON body to the requestBody struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Log the ID and data to be updated
	log.Printf("Updating product with ID: %d, Name: %s, Description: %s, Price: %f",
		requestBody.ID, requestBody.Name, requestBody.Description, requestBody.Price)

	// Update the product in the database
	result, err := config.DB.Exec("UPDATE Products SET name = ?, description = ?, price = ? WHERE ID = ?",
		requestBody.Name, requestBody.Description, requestBody.Price, requestBody.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found "})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {
	var requestBody struct {
		ID int `json:"id"`
	}

	// Bind the JSON body to the requestBody struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Use the ID to delete the product
	result, err := config.DB.Exec("DELETE FROM Products WHERE id = ?", requestBody.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
