package session

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/setis-project/api/pkg/database"
)

func SetRoutes(router *gin.RouterGroup, db *database.Db, redisCli *redis.Client) {
	path := "/session"
	router.GET(path+"/register", Refresh(redisCli))
}
