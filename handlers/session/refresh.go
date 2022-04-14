package session

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/setis-project/api/controllers"
)

// method: GET
func Refresh(redisCli *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, err := ctx.Request.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		refresh, err := ctx.Request.Cookie("refresh-session")
		if err != nil {
			if err == http.ErrNoCookie {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		sessionToken, err := uuid.Parse(session.Value)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not parse refresh token"})
			return
		}
		refreshToken, err := uuid.Parse(refresh.Value)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not parse refresh token"})
			return
		}

		newSession, err := controllers.RefreshSession(redisCli, sessionToken, refreshToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// fiz maxAge
		ctx.SetCookie(
			"session",
			newSession.Token.String(),
			int(time.Since(newSession.Expiry).Milliseconds()),
			"/session",
			os.Getenv("SERVER_DOMAIN"),
			true,
			true,
		)
		ctx.SetCookie(
			"refresh-session",
			newSession.RefreshToken.Token.String(),
			int(time.Since(newSession.RefreshToken.Expiry).Milliseconds()),
			"/session",
			os.Getenv("SERVER_DOMAIN"),
			true,
			true,
		)
		ctx.Status(http.StatusOK)
	}
}
