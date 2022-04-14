package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/setis-project/api/core"
)

func EnsureAuthToken(redisCli *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionCookie, err := ctx.Request.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		sessionToken, err := uuid.Parse(sessionCookie.Value)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid value for session token"})
			return
		}
		session, err := core.GetSession(redisCli, sessionToken)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if session.Expiry.Before(time.Now()) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
