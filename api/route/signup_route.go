package route

import (
	"time"

	bootstrap "github.com/FxIvan/go-backend/boostrap"
	"github.com/FxIvan/go-backend/mongo"
	"github.com/gin-gonic/gin"
)

func NewSignupRouter(router *gin.RouterGroup, env *bootstrap.Env, timeout time.Duration, db mongo.Database) {

}
