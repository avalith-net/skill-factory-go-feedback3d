package router

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-feedback3d/controller"
	_ "github.com/avalith-net/skill-factory-go-feedback3d/docs"
	"github.com/avalith-net/skill-factory-go-feedback3d/middleware"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
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
	r.Use(cors.AllowAll())

	r.LoadHTMLGlob("templates/*")

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
			jwt.POST("/fbRequest", controller.RequestFeedback)
			jwt.PATCH("/users/edit", controller.EditUser)
			jwt.GET("/dashboard", controller.GetDashboard)
			jwt.GET("/users/search", controller.GetByFullName)
			jwt.GET("/users/get/:id", controller.GetGeneralProfile)
			jwt.POST("/changePassword", controller.ChangePassEmail)
			jwt.GET("/selectedFeedback", controller.GetSelectedFeedback)
			jwt.PATCH("/users/report/:feed_id", controller.ReportFeed)
			jwt.GET("/users/get_feedback_sent/:id", controller.GetFeedsSent)
			jwt.GET("/users/my_feedbacks", controller.GetAllMyFeedbacks)
			jwt.GET("/users/get_feed_requests_from_me", controller.GetFeedRequestsFromMe)

			admin := jwt.Group("/")
			admin.Use(middleware.IsAdmin())
			{
				admin.PATCH("/users/ban/:id", controller.BanUser)
				admin.PATCH("/users/feed_state", controller.FeedbackState)
			}
		}
		endpoints.GET("/feedback", func(c *gin.Context) {
			c.HTML(http.StatusOK, "feedback.tmpl", gin.H{})
		})
		endpoints.GET("/changePassword", func(c *gin.Context) {
			c.HTML(http.StatusOK, "change_pass.tmpl", gin.H{})
		})
		endpoints.POST("/recoverPass", controller.RecoverPass)
	}

	//-----------------------------------------------------------------------------------------------
	// use ginSwagger middleware to serve the API docs
	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run()
}
