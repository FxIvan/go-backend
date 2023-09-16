package route

import (
	"fmt"
	"time"

	bootstrap "github.com/FxIvan/go-backend/boostrap"
	"github.com/FxIvan/go-backend/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	fmt.Println("Setup route")
	fmt.Print("Env: ", env)
	fmt.Print("Timeout: ", timeout)
	fmt.Print("DB: ", db)
	fmt.Print("Gin: ", gin)
}
