package admin

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/setis-project/api/controllers"
	"github.com/setis-project/api/core"
	"github.com/setis-project/api/core/models"
	"github.com/setis-project/api/utils"
)

// Admin login
// @Summary      admin login
// @Description  login an admin
// @Tags         Admin
//
// @Accept       json
// @Param        email     body  string  true  "account email"
// @Param        password  body  string  true  "account password"
//
// @Success      200
// @Failure      400  {object}  models.ApiRequestErrors "Error on request fields"
// @Failure      400  {object}  models.ApiError "Execution error"
// @Failure      404  {object}  models.ApiError "Execution error"
//
// @Router       /v1/admin/account/login [post]
// @Security     securitydefinitions.apikey
func Login(db *core.Db, redisCli *redis.Client) gin.HandlerFunc {
	type request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	return func(ctx *gin.Context) {
		body := request{}
		if err := ctx.BindJSON(&body); err != nil {
			if out, ok := utils.GetBindErrors(err); ok {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ApiRequestErrors{Errors: out})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ApiError{Error: err.Error()})
			return
		}
		session, err := controllers.Login(db, redisCli, body.Email, body.Password)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, models.ApiError{Error: err.Error()})
			return
		}

		// fix maxAge
		ctx.SetCookie(
			"session",
			session.Token.String(),
			int(time.Since(session.Expiry).Milliseconds()),
			"/admin",
			os.Getenv("SERVER_DOMAIN"),
			true,
			true,
		)
		ctx.SetCookie(
			"refresh-session",
			session.RefreshToken.Token.String(),
			int(time.Since(session.RefreshToken.Expiry).Milliseconds()),
			"/admin",
			os.Getenv("SERVER_DOMAIN"),
			true,
			true,
		)
		ctx.Status(http.StatusOK)
	}
}
