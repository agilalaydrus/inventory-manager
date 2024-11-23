package models

import (
	"database/sql"
)

// Order represents the structure of an order in the database
type Order struct {
	OrderID   int          `json:"order_id"`
	ProductID int          `json:"product_id"`
	Quantity  int          `json:"quantity"`
	OrderDate sql.NullTime `json:"order_date"`
}
