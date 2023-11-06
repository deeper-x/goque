package goque

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Check connection parameters
func Check() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println("hello, handler")
		c.Next()
	}
}
