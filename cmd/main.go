package main

import (
	"fmt"
	"time"

	route "github.com/FxIvan/go-backend/api/route"
	bootstrap "github.com/FxIvan/go-backend/boostrap"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	app := bootstrap.App()
	env := app.Env

	db := app.Mongo.Database(env.DBName)

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
