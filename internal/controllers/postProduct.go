package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postProduct struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"omitempty,max=250"`
}

type Product struct {
	GUID        string  `json:"guid"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdAt"`
}

func PostProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload postProduct
		var ctx = c.Request.Context()

		// Bind JSON
		if e := c.ShouldBindJSON(&payload); e != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": e.Error(),
			})
			return
		}
		// Insert product
		var guid = uuid.New().String()
		var createdAt = time.Now().Format(time.RFC3339)
		if _, e := db.ExecContext(ctx, "INSERT INTO products (guid,name,price,description,createdAt) VALUES (?,?,?,?,?)", guid, payload.Name, payload.Price, payload.Description, createdAt); e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": e.Error(),
			})
			return
		}
		// Fetch created product
		var product Product
		var row = db.QueryRowContext(ctx, "SELECT guid,name,price,description,createdAt FROM products WHERE guid=?", guid)
		if e := row.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": e.Error(),
			})
			return
		}
		// Write GUID to response header
		c.Writer.Header().Add("Location", fmt.Sprintf("/products/%s", guid))
		c.JSON(http.StatusCreated, product)
	}
}
