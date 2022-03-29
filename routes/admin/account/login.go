package account

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/setis-project/api/pkg"
	repo "github.com/setis-project/api/repo/admin/account"
	"github.com/setis-project/api/utils"
)

// method: POST
func Login(db *pkg.Db, redis *redis.Client) gin.HandlerFunc {
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
		credentials, err := repo.Login(db, redis, body.Email, body.Password)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		expiryDuration := time.Minute * 5
		token := uuid.NewString()
		redis.Set(token, credentials.Id, expiryDuration)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		ctx.SetCookie(
			"session",
			token,
			int(expiryDuration.Milliseconds()),
			"/admin",
			os.Getenv("DOMAIN"),
			true,
			true,
		)
		ctx.Status(http.StatusOK)
	}
}
