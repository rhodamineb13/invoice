package middleware

import (
	"invoice/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *helper.CustomError:
				c.AbortWithStatusJSON(e.Status, gin.H{
					"error": e.Message,
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			}
		}
	}
}
