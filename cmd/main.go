package main

import (
	"fmt"
	"time"

	bootstrap "github.com/FxIvan/go-backend/boostrap"
	route "github.com/FxIvan/go-backend/boostrap/api/route"
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

	//gin.Run(fmt.Sprintf(":%d", env.SERVER_ADDRESS))
}
