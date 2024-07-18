package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiboquinha/curdgo/src/configuration/erros"
	"github.com/xiboquinha/curdgo/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		erro := erros.NewBadRequestErr(
			fmt.Sprintf("preencheu algo errado, error=%s\n", err.Error()))
		c.JSON(erro.Code, erro)
		return
	}

	fmt.Println(userRequest)
}
