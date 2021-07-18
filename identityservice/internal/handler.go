package internal

import (
	"github.com/bariis/microservices-demo/client"
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
