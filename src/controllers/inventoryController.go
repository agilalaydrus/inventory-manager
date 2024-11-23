package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inventory-management/src/config"
	"inventory-management/src/models"
	"net/http"
)

// GetInventory retrieves the inventory for a specific product by its ID from the JSON body
func GetInventory(c *gin.Context) {
	var request struct {
		ProductID string `json:"product_id"`
	}

	// Bind the JSON input to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	var inventory models.Inventory

	// Query the inventory from the database
	err := config.DB.QueryRow("SELECT product_id, quantity, location, min_stock, max_stock FROM Inventory WHERE product_id = ?", request.ProductID).Scan(&inventory.ProductID, &inventory.Quantity, &inventory.Location, &inventory.MinStock, &inventory.MaxStock)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// CreateInventory creates a new inventory record for a specific product from the JSON body
func CreateInventory(c *gin.Context) {
	var request struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
		Location  string `json:"location"`
		MinStock  int    `json:"min_stock"`
		MaxStock  int    `json:"max_stock"`
	}

	// Bind the JSON input to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Insert the new inventory record into the database
	_, err := config.DB.Exec("INSERT INTO Inventory (product_id, quantity, location, min_stock, max_stock) VALUES (?, ?, ?, ?, ?)", request.ProductID, request.Quantity, request.Location, request.MinStock, request.MaxStock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Inventory created successfully"})
}

// UpdateInventory updates the inventory for a specific product from the JSON body
func UpdateInventory(c *gin.Context) {
	var request struct {
		ProductID      string `json:"product_id"`
		ChangeQuantity int    `json:"change_quantity"`
		Location       string `json:"location"`
	}

	// Bind the JSON input to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Retrieve the current inventory
	var currentInventory models.Inventory
	err := config.DB.QueryRow("SELECT product_id, quantity, location, min_stock, max_stock FROM Inventory WHERE product_id = ?", request.ProductID).Scan(&currentInventory.ProductID, &currentInventory.Quantity, &currentInventory.Location, &currentInventory.MinStock, &currentInventory.MaxStock)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	// Update the quantity based on the change quantity
	newQuantity := currentInventory.Quantity + request.ChangeQuantity

	// Ensure the new quantity does not go below zero
	if newQuantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity cannot be negative"})
		return
	}

	// Update the inventory record in the database
	_, err = config.DB.Exec("UPDATE Inventory SET quantity = ?, location = ? WHERE product_id = ?", newQuantity, request.Location, request.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
}

// DeleteInventory deletes the inventory record for a specific product from the JSON body
func DeleteInventory(c *gin.Context) {
	var request struct {
		ProductID string `json:"product_id"`
	}

	// Bind the JSON input to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Delete the inventory record from the database
	_, err := config.DB.Exec("DELETE FROM Inventory WHERE product_id = ?", request.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory deleted successfully"})
}

// ReduceQuantity mengurangi quantity dari inventory berdasarkan ProductID
func ReduceQuantity(c *gin.Context) {
	var request struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}

	// Bind the JSON input to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Log request untuk debugging
	fmt.Printf("Received request: %+v\n", request)

	// Pastikan jumlah yang ingin dikurangi lebih besar dari nol
	if request.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than zero"})
		return
	}

	// Retrieve the current inventory
	var currentInventory models.Inventory
	err := config.DB.QueryRow("SELECT product_id, quantity, location, min_stock, max_stock FROM Inventory WHERE product_id = ?", request.ProductID).Scan(&currentInventory.ProductID, &currentInventory.Quantity, &currentInventory.Location, &currentInventory.MinStock, &currentInventory.MaxStock)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	// Ensure the quantity to reduce is valid
	if request.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than zero"})
		return
	}

	// Calculate new quantity
	newQuantity := currentInventory.Quantity - request.Quantity

	// Ensure the new quantity does not go below zero
	if newQuantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient quantity"})
		return
	}

	// Update the inventory record in the database
	_, err = config.DB.Exec("UPDATE Inventory SET quantity = ? WHERE product_id = ?", newQuantity, request.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quantity reduced successfully", "new_quantity": newQuantity})
}
