package middleware

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func CORSMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		method := string(c.Request.Method())
		c.Header("Access-Control-Allow-Origin", string(c.GetHeader("Origin")))
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", string(c.GetHeader("Access-Control-Request-Method")))
			c.Header("Access-Control-Allow-Headers", string(c.GetHeader("Access-Control-Request-Headers")))
			c.Header("Access-Control-Max-Age", "7200")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next(ctx)
	}
}
