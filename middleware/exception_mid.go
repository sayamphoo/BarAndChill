package middleware

import (
	"sayamphoo/microservice/models/domain"

	"github.com/gin-gonic/gin"
)

func ExceptionHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(c *gin.Context) {

			if r := recover(); r != nil {
				data, ok := r.(domain.UtilityModel)
				if !ok {
					return
				}

				c.JSON(data.Code, data)
				c.Abort()
				return
			}
		}(c)
		c.Next()
	}
}
