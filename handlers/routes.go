package handlers

import (
	"github.com/blotin1993/feedback-api/middleware"
	"github.com/blotin1993/feedback-api/routers"
	"github.com/gin-gonic/gin"
)

//SetRoutes  ...
func SetRoutes() {
	r := gin.Default()
	r.Use(middleware.CheckDb())

	//Endpoints ------------------------------------------------------------------------------------
	r.POST("/sign_up", routers.SignUp)
	r.POST("/login", routers.Login)
	r.POST("/feedback", middleware.ValidateJWT(), routers.FeedbackTry)
	r.POST("/setProfilePic", middleware.ValidateJWT(), routers.SetProfilePicture)
	r.POST("/recoverPass", middleware.ValidateJWT(), routers.RecoverPass)
	r.GET("/getfb", middleware.ValidateJWT(), routers.GetFeed)
	r.POST("/changePassword", routers.ChangePassEmail)
	//-----------------------------------------------------------------------------------------------

	r.Run()
}
