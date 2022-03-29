package admin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func EnsureAuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if cookie.Expires.Before(time.Now()) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
