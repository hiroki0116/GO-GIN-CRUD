package controllers

import (
	"database/sql"
	"go-gin-review/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rows *sql.Rows
		var products []Product
		var e error

		if rows, e = db.Query("SELECT guid,name,price,description,createdAt FROM products"); e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var product Product
			if e := rows.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); e != nil {
				var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
				c.JSON(http.StatusInternalServerError, res)
				return
			}
			products = append(products, product)
		}
		// 404 if no rows found
		if len(products) == 0 {
			var res = internal.NewHTTPResponse(http.StatusNotFound, e)
			c.JSON(http.StatusNotFound, res)
			return
		}

		var res = internal.NewHTTPResponse(http.StatusOK, products)
		c.JSON(http.StatusOK, res)
	}
}
