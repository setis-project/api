package session

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/setis-project/api/core"
)

func SetRoutes(router *gin.RouterGroup, db *core.Db, redisCli *redis.Client) {
	path := "/session"
	router.GET(path+"/refresh", Refresh(redisCli))
}
