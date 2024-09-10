package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var origins = []string{
	"http://localhost:5173",
	"http://localhost:3000",
}

func Cores() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("origin")

		for _, o := range origins {
			if o == origin {
				c.Header("Access-Control-Allow-Origin", o)
				c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
				c.Header("Access-Control-Allow-Credentials", "true")
				c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
				if c.Request.Method == "OPTIONS" {
					c.JSON(http.StatusOK, "")
					c.AbortWithStatus(204)
					return
				}
				c.Next()
			}
		}
	}
}
