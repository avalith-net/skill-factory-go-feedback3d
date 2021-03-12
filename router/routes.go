package router

import (
	"github.com/blotin1993/feedback-api/controller"
	_ "github.com/blotin1993/feedback-api/docs"
	"github.com/blotin1993/feedback-api/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title feedback API
// @version 1.0
// @description Aplicación que permite realizar feedbacks entre los miembros de un equipo de trabajo.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email vlotin_gaming@gmail.com

// @license.name Avalith

// @host localhost:8080
// @BasePath /
// @schemes http https
func SetRoutes() {
	//set router
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
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
