package router

import (
	"net/http"
	"os"

	"github.com/blotin1993/feedback-api/controller"
	_ "github.com/blotin1993/feedback-api/docs"
	"github.com/blotin1993/feedback-api/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title feedback API
// @version 1.0
// @description Application used to give feedback between members of a workgroup.
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

	endpoints := r.Group("/")
	//Endpoints ------------------------------CHEQUEAR GET POST , ETC. GET OBTENER, POST CREAR,ETC.------------------------------------------------------
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
			jwt.POST("/fbRequest", controller.RequestFeedback)
		}
		endpoints.GET("/feedback", func(c *gin.Context) {
			c.HTML(http.StatusOK, "feedback.tmpl", gin.H{})
		})
		endpoints.GET("/changePassword", func(c *gin.Context) {
			c.HTML(http.StatusOK, "change_pass.tmpl", gin.H{})
		})
		endpoints.POST("/recoverPass", controller.RecoverPass)
		endpoints.POST("/changePassword", controller.ChangePassEmail)
	}

	// Admin routes
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": os.Getenv("ADMIN_PASS"),
	}))
	authorized.GET("/users", controller.GetUsers)
	authorized.PATCH("/users/:email", controller.BanUser)

	//-----------------------------------------------------------------------------------------------
	// use ginSwagger middleware to serve the API docs
	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run()
}
