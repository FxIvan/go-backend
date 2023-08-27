package main

import (
	"fmt"

	"github.com/FxIvan/go-backend/bootstrap"
)

func main() {
	fmt.Println("Hello, World!")
	app := bootstrap.App()

	fmt.Println(app.Env)
}
