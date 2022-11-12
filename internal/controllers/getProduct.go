package controllers

import (
	"database/sql"
	"go-gin-review/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

type guidBinding struct {
	GUID string `uri:"guid" binding:"required,uuid4"`
}

func GetProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product Product
		var binding guidBinding
		var ctx = c.Request.Context()
		if e := c.ShouldBindUri(&binding); e != nil {
			var res = internal.NewHTTPResponse(http.StatusBadRequest, e)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		var row = db.QueryRowContext(ctx, "SELECT guid,name,price,description,createdAt FROM products WHERE guid = ?", binding.GUID)
		if e := row.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); e != nil {
			if e == sql.ErrNoRows {
				var res = internal.NewHTTPResponse(http.StatusNotFound, e)
				c.JSON(http.StatusNotFound, res)
				return
			}

			var res = internal.NewHTTPResponse(http.StatusBadRequest, e)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		var res = internal.NewHTTPResponse(http.StatusOK, product)
		c.JSON(http.StatusOK, res)
	}
}
