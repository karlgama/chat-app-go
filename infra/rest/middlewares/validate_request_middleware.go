package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO not working
func ValidateRequest[T any](next func(*gin.Context, *T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := new(T)
		if err := c.ShouldBindJSON(params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		next(c, params)
	}
}
