package main

import (
	"fmt"

	bootstrap "github.com/FxIvan/go-backend/boostrap"
)

func main() {
	fmt.Println("Hello, World!")
	app := bootstrap.App()
	env := app.Env

	app.Mongo.Database(env.DBName)
}
