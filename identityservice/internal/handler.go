package internal

import (
	"github.com/bariis/microservices-demo/client"
	"github.com/bariis/microservices-demo/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *IdentityService) SignUpHandler(context *gin.Context) {
	var input *client.Client
	context.BindJSON(&input)

	user, err := SignUp(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}

	context.JSON(http.StatusOK, user)
}

func (i *IdentityService) SignInHandler(context *gin.Context) {
	var input *client.Client
	context.BindJSON(&input)

	user, err := SignIn(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}

	token := jwt.GenerateToken(user.ID)

	context.JSON(http.StatusOK, jwt.TokenResponse{AccessToken: token})
}