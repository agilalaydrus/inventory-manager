package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"inventory-management/src/config"
	"inventory-management/src/models"
	"net/http"
	"strconv"
	"time"
)

func formatNullTime(nullTime sql.NullTime) interface{} {
	if nullTime.Valid {
		return nullTime.Time
	}
	return nil
}

// CreateOrder creates a new order
func CreateOrder(c *gin.Context) {
	var order models.Order
	// Bind JSON input to Order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Check if OrderDate is null or not provided, and set it to current time
	if !order.OrderDate.Valid {
		order.OrderDate = sql.NullTime{Time: time.Now(), Valid: true}
	}

	// Insert the order into the database
	result, err := config.DB.Exec("INSERT INTO Orders (product_id, quantity, order_date) VALUES (?, ?, ?)", order.ProductID, order.Quantity, order.OrderDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order: " + err.Error()})
		return
	}

	// Get the last inserted ID
	orderID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order ID: " + err.Error()})
		return
	}

	// Set the OrderID in the response
	order.OrderID = int(orderID)
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order_id": order.OrderID})
}

func GetOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order

	// Query the order from the database
	query := "SELECT order_id, product_id, quantity, order_date FROM Orders WHERE order_id = ?"
	err = config.DB.QueryRow(query, id).Scan(&order.OrderID, &order.ProductID, &order.Quantity, &order.OrderDate)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		} else {
			// Log the error with additional context
			fmt.Printf("Error querying order with ID %d: %v\n", id, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	// Convert NullTime to regular time if valid
	if order.OrderDate.Valid {
		c.JSON(http.StatusOK, gin.H{
			"order_id":   order.OrderID,
			"product_id": order.ProductID,
			"quantity":   order.Quantity,
			"order_date": order.OrderDate.Time.Format("2006-01-02 15:04:05"),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"order_id":   order.OrderID,
			"product_id": order.ProductID,
			"quantity":   order.Quantity,
			"order_date": nil,
		})
	}
}
