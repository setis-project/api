package account

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/setis-project/api/pkg/database"
	repo "github.com/setis-project/api/repo/admin/account"
	"github.com/setis-project/api/utils"
)

// method: POST
func Login(db *database.Db, redisCli *redis.Client) gin.HandlerFunc {
	type request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	return func(ctx *gin.Context) {
		body := request{}
		if err := ctx.BindJSON(&body); err != nil {
			if out, ok := utils.GetBindErrors(err); ok {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		session, err := repo.Login(db, redisCli, body.Email, body.Password)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		// fix maxAge
		ctx.SetCookie(
			"session",
			session.Token.String(),
			int(time.Since(session.Expiry).Milliseconds()),
			"/admin",
			os.Getenv("DOMAIN"),
			true,
			true,
		)
		ctx.SetCookie(
			"refresh-session",
			session.RefreshToken.Token.String(),
			int(time.Since(session.RefreshToken.Expiry).Milliseconds()),
			"/admin",
			os.Getenv("DOMAIN"),
			true,
			true,
		)
		ctx.Status(http.StatusOK)
	}
}
