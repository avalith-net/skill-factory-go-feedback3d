package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/blotin1993/feedback-api/middleware"
	"github.com/blotin1993/feedback-api/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//SetRoutes  ...
func SetRoutes() {
	router := mux.NewRouter()

	//Endpoints ------------------------------------------------------------------------------------
	router.HandleFunc("/sign_up", middleware.CheckDb(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDb(routers.Login)).Methods("POST")
	router.HandleFunc("/feedback", middleware.CheckDb(middleware.ValidateJWT(routers.FeedbackTry))).Methods("POST")
	router.HandleFunc("/setProfilePic", middleware.CheckDb(middleware.ValidateJWT(routers.SetProfilePicture))).Methods("POST")
	//-----------------------------------------------------------------------------------------------

	PORT := os.Getenv("PORT")
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
