package router

import (
	"github.com/blotin1993/feedback-api/controller"
	"github.com/blotin1993/feedback-api/middleware"
	"github.com/gin-gonic/gin"
)

//SetRoutes  ...
func SetRoutes() {
	r := gin.Default()
	r.Use(middleware.CheckDb())

	//Endpoints ------------------------------------------------------------------------------------
	r.POST("/sign_up", controller.SignUp)
	r.POST("/login", controller.Login)
	r.POST("/feedback", middleware.ValidateJWT(), controller.FeedbackTry)
	r.POST("/setProfilePic", middleware.ValidateJWT(), controller.SetProfilePicture)
	r.POST("/recoverPass", middleware.ValidateJWT(), controller.RecoverPass)
	r.GET("/getfb", middleware.ValidateJWT(), controller.GetFeed)
	r.POST("/changePassword", controller.ChangePassEmail)
	//-----------------------------------------------------------------------------------------------

	r.Run()
}
