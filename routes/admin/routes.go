package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/setis-project/api/pkg/database"
	"github.com/setis-project/api/routes/admin/account"
)

func SetRoutes(router *gin.RouterGroup, db *database.Db, redis *redis.Client) {
	path := "/admin"
	group := router.Group(path)
	account.SetRoutes(group, db, redis)
}
