package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	ginContext := "GinContext"
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginContext, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
