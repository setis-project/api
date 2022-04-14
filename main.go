package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/setis-project/api/core"
	"github.com/setis-project/api/docs"
	"github.com/setis-project/api/handlers"
)

// @title          Setis Project API
// @description    This is the Setis Project's API.
// @BasePath       /v1
//
// @contact.name   API Support
// @contact.email  setisproject@gmail.com
//
// @license.name   GPL-3.0 License
// @license.url    https://github.com/setis-project/api/blob/main/LICENSE

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	docs.SwaggerInfo.Version = os.Getenv("API_VERSION")
	docs.SwaggerInfo.Host = os.Getenv("SERVER_DOMAIN")

	ctx := context.Background()
	db, err := core.DbConnect(&ctx)
	if err != nil {
		log.Fatal(err)
	}

	redisCli, err := core.RedisConnect()
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.New()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handlers.SetRoutes(engine, db, redisCli)
	engine.Run(os.Getenv("SERVER_DOMAIN"))
}
