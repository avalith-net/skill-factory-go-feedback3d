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
	endpoints := r.Group("/")
	//Endpoints ------------------------------------------------------------------------------------
	endpoints.Use(middleware.CheckDb())
	{
		endpoints.POST("/sign_up", controller.SignUp)
		endpoints.POST("/login", controller.Login)

		//using jwt
		jwt := endpoints.Group("/")
		jwt.Use(middleware.ValidateJWT())
		{
			jwt.POST("/feedback", controller.FeedbackTry)
			jwt.POST("/setProfilePic", controller.SetProfilePicture)
			jwt.POST("/recoverPass", controller.RecoverPass)
			jwt.GET("/getfb", controller.GetFeed)
			jwt.POST("/fbRequest", controller.RequestFeedback)
		}

		endpoints.POST("/changePassword", controller.ChangePassEmail)
	}
	//-----------------------------------------------------------------------------------------------
	// use ginSwagger middleware to serve the API docs
	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run()
}
