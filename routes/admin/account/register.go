package account

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/setis-project/api/pkg/database"
	repo "github.com/setis-project/api/repo/admin/account"
	"github.com/setis-project/api/utils"
)

// method: POST
func Register(db *database.Db) gin.HandlerFunc {
	type request struct {
		FirstName string `json:"name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required"`
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
		err := repo.Register(
			db,
			body.FirstName,
			body.LastName,
			body.Email,
			body.Password,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Status(http.StatusCreated)
	}
}
