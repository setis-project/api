package account

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	mw "github.com/setis-project/api/middlewares/admin"
	"github.com/setis-project/api/pkg/database"
)

func SetRoutes(router *gin.RouterGroup, db *database.Db, redis *redis.Client) {
	path := "/account"
	router.POST(path+"/register", mw.EnsureAuthToken(), Register(db))
	router.POST(path+"/login", Login(db, redis))
}
