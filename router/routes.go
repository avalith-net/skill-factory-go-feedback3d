package router

import (
	"net/http"

	"github.com/blotin1993/feedback-api/controller"
	_ "github.com/blotin1993/feedback-api/docs"
	"github.com/blotin1993/feedback-api/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title feedback API
// @version 1.0
// @description Aplicación que permite realizar feedbacks entre los miembros de un equipo de trabajo. Es un proceso mediante el cual se recogen los comentarios de los miembros del equipo (desarrolladores), de los jefes de equipo, de los clientes internos, de los clientes externos o de otras partes interesadas(que los suben directamente o los realizan a través del equipo de entrega), así como una autoevaluación de los miembros de cada equipo.
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
	r.LoadHTMLGlob("templates/*")

	// Front end>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	r.GET("/changePassword", func(c *gin.Context) {
		c.HTML(http.StatusOK, "change_pass.tmpl", gin.H{})
	})
	endpoints := r.Group("/")
	//Endpoints Backend------------------------------------------------------------------------------------
	endpoints.Use(middleware.CheckDb())
	{
		r.POST("/sign_up", controller.SignUp)
		r.POST("/login", controller.Login)

		//using jwt
		jwt := r.Group("/")
		jwt.Use(middleware.ValidateJWT())
		{
			r.POST("/feedback", controller.FeedbackAttempt)
			r.POST("/setProfilePic", controller.SetProfilePicture)
			r.POST("/recoverPass", controller.RecoverPass)
			r.POST("/fbRequest", controller.RequestFeedback)
		}
		r.POST("/changePassword", controller.ChangePassEmail)
	}
	//-----------------------------------------------------------------------------------------------
	// use ginSwagger middleware to serve the API docs
	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run()
}
