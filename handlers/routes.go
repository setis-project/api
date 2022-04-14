package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/setis-project/api/core"
	"github.com/setis-project/api/handlers/admin"
	"github.com/setis-project/api/handlers/session"
)

func SetRoutes(engine *gin.Engine, db *core.Db, redisCli *redis.Client) {
	group := engine.Group("/v1")
	admin.SetRoutes(group, db, redisCli)
	session.SetRoutes(group, db, redisCli)
}
