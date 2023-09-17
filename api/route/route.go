package route

import (
	"time"

	bootstrap "github.com/FxIvan/go-backend/boostrap"
	"github.com/FxIvan/go-backend/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("") //generador de ruta, podriamos poner gin.Group("/api/v1") para que todas las rutas tengan ese prefijo
	NewSignupRouter(publicRouter, env, timeout, db)
}
