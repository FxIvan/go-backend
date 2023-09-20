package route

import (
	"fmt"
	"time"

	"github.com/FxIvan/go-backend/api/controller"
	bootstrap "github.com/FxIvan/go-backend/boostrap"
	"github.com/FxIvan/go-backend/domain"
	"github.com/FxIvan/go-backend/mongo"
	"github.com/FxIvan/go-backend/repository"
	"github.com/FxIvan/go-backend/usecase"
	"github.com/gin-gonic/gin"
)

func NewSignupRouter(router *gin.RouterGroup, env *bootstrap.Env, timeout time.Duration, db mongo.Database) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	fmt.Print(ur)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignUsecase(ur, timeout),
		Env:           env,
	}

	router.POST("/signup", sc.Signup)

}
