package controller

import (
	"fmt"
	"net/http"

	bootstrap "github.com/FxIvan/go-backend/boostrap"
	"github.com/FxIvan/go-backend/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	fmt.Println(request)
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Something went wrong"})
		return
	}

	c.JSON(http.StatusCreated, user)

}
