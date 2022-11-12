package controllers

import (
	"database/sql"
	"go-gin-review/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var guid guidBinding
		var ctx = c.Request.Context()
		if e := c.ShouldBindUri(&guid); e != nil {
			var res = internal.NewHTTPResponse(http.StatusBadRequest, e)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		var result sql.Result
		var e error
		if result, e = db.ExecContext(ctx, "DELETE FROM products WHERE guid = ?", guid.GUID); e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		if nProducts, _ := result.RowsAffected(); nProducts == 0 {
			var res = internal.NewHTTPResponse(http.StatusNotFound, e)
			c.JSON(http.StatusNotFound, res)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
