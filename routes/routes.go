package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/setis-project/api/pkg/database"
	"github.com/setis-project/api/routes/admin"
	"github.com/setis-project/api/routes/session"
)

func SetRoutes(engine *gin.Engine, db *database.Db, redisCli *redis.Client) {
	group := engine.Group("")
	admin.SetRoutes(group, db, redisCli)
	session.SetRoutes(group, db, redisCli)
}
