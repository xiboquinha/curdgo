package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xiboquinha/curdgo/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserByID/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
