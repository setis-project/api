package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	mw "github.com/setis-project/api/controllers/middlewares"
	"github.com/setis-project/api/core"
)

func SetRoutes(router *gin.RouterGroup, db *core.Db, redisCli *redis.Client) {
	group := router.Group("/admin")

	account := group.Group("/account")
	account.POST("/register", mw.EnsureAuthToken(redisCli), Register(db))
	account.POST("/login", Login(db, redisCli))
}
