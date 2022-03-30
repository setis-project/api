package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/setis-project/api/pkg/database"
	predis "github.com/setis-project/api/pkg/redis"
	"github.com/setis-project/api/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	db, err := database.Connect(&ctx)
	if err != nil {
		log.Fatal(err)
	}

	redisCli, err := predis.Connect()
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.New()
	routes.SetRoutes(engine, db, redisCli)
	engine.Run(":8080")
}
