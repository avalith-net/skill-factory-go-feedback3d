package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/db"
	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	services "github.com/JoaoPaulo87/skill-factory-go-feedback3d/services/email"
)

const (
	timeFormat = "2006.01.02 15:04:05"
)

// RequestFeedback godoc
// @Description get string by id
// @id RequestFeedback
// @Summary is used to request a feedback to other user.
// @Param id query string true "Account ID"
// @Param Authorization header string true "jwt token"
// @Success 201 {string} string "Email sended successfully."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "internal error"
// @Failure 500 {string} string "An error has ocurred sending the email."
// @Failure default {string} string "An error has ocurred"
// @Router /fbRequest [post]
func RequestFeedback(c *gin.Context) {
	id := c.Query("id")
	if len(id) < 1 || id == IDUser {
		c.String(http.StatusBadRequest, "Error with the request.")
		return
	}

	user, err := db.GetUser(id)
	if err != nil {
		c.String(http.StatusBadRequest, "User not found.")
		return
	}

	loggedUser, err := db.GetUser(IDUser)
	if err != nil {
		c.String(http.StatusBadRequest, "User does not exists.")
		return
	}
	//------------------------------------------------------------------------------------------

	if !user.Enabled {
		c.String(http.StatusBadRequest, "Cannot ask feedback, the target user is banned")
		return
	}

	// Como yo, la persona logeada, soy el que pido feedback, le tengo que pasar mis datos y
	//agregarselos al otro usuario en el userAsksfeed. En feedrequest tengo que agregarme los
	//datos del otro usuario a mi. Y luego agregar ambos con modify al usuario.

	// PASO A SEGUIR: Una vez que lo creo, ANTES DE AGREGARLO AL USUARIO, lo que voy a hacer es
	//con feedRequestedID busco el objeto y ese es el que tengo que agregar al Usuario para
	//que se me agrege con el fking ObjId y asi poder borrarlo!!!!
	var feedRequested models.FeedbacksRequested

	feedRequested.RequestedUserID = id
	feedRequested.UserLoggedID = IDUser
	feedRequested.RequestedUserName = user.Name
	feedRequested.RequestedUserLastName = user.LastName
	feedRequested.SentDate = time.Now()

	feedRequestedID, isRequestCreated, err := db.AddFeedbackRequested(feedRequested)
	if err != nil {
		c.String(http.StatusBadRequest, "Cannot create a feedback request to target user in feedback request.")
		return
	}
	if !isRequestCreated {
		c.String(http.StatusBadRequest, "Something went wrong trying create a request to target user in feedback request.")
		return
	}

	fmt.Println(feedRequestedID)
	// feedRequestedObj, err := db.GetSelectedFeedBackRequestObj(feedRequestedID)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Cannot get objID of feedbackRequest in feedback request.")
	// 	return
	// }

	var userAskingFeed models.UsersAskedFeed

	userAskingFeed.UserWhoAskFeedID = IDUser
	userAskingFeed.UserAskedForFeedID = id
	userAskingFeed.NameWhoAskFeed = loggedUser.Name
	userAskingFeed.LastNameWhoAskFeed = loggedUser.LastName
	userAskingFeed.SentDate = time.Now()

	userAskingFeedID, isRequestCreated, err := db.AddUsersAsksFeed(userAskingFeed)
	if err != nil {
		c.String(http.StatusBadRequest, "Cannot create a userAsksFeed instance in feedback request.")
		return
	}
	if !isRequestCreated {
		c.String(http.StatusBadRequest, "Something went wrong trying create a userAsksFeed instance in feedback request. ID: "+userAskingFeedID)
		return
	}

	//Ahora que tengo el ID del userAskingFeedID, tengo que persistir el userAskingFeedObj en el usuario. Todo esto para luego en feedback_attempt buscar el obj con el ID
	// que esta relacionado al usuario y poder borrarlo tambien, porque ahora borra el obj de la bd pero no se actualiza en el usuario ese borrado.
	// userAskingFeedObj, err := db.GetUserAskingForFeedBackObj(userAskingFeedID)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Cannot get objID of userAskingFeed in feedback request.")
	// 	return
	// }

	var modifiedLoggedUser models.User

	copier.Copy(&modifiedLoggedUser, &loggedUser)

	//modifiedLoggedUser.FeedbackStatus.FeedbacksRequested = feedRequestedObj

	isLoggedUserModified, err := db.ModifyUser(modifiedLoggedUser, IDUser)
	if err != nil {
		c.String(http.StatusBadRequest, "Cannot modify the logged user.")
		return
	}
	if !isLoggedUserModified {
		c.String(http.StatusBadRequest, "Something went wrong trying modify the logged user.")
		return
	}
	//------------------------------------------------------------------------------------------

	var modifiedTargetUser models.User

	copier.Copy(&modifiedTargetUser, &user)

	if !modifiedTargetUser.Enabled {
		c.String(http.StatusBadRequest, "Cannot ask feedback, the logged user is banned")
		return
	}
	//modifiedTargetUser.FeedbackStatus.UsersAskedFeed = userAskingFeedObj

	isTargetUserModified, err := db.ModifyUser(modifiedTargetUser, id)
	if err != nil {
		c.String(http.StatusBadRequest, "Cannot modify the target user.")
		return
	}

	if !isTargetUserModified {
		c.String(http.StatusBadRequest, "Something went wrong trying modify the target user.")
		return
	}
	//------------------------------------------------------------------------------------------
	// An email is send to all admins with the report notification.
	bodyString := "Hi <b><i>" + user.Name + "</i></b>!\n" +
		"I'd like to ask a few questions about your working experience with me. It's important to help me to improve." +
		"Follow this link to give me feedback: <b><i>http:localhost:8080/feedback?target_id=" + IDUser +
		"</i></b>\n<br> Thanks for your time!\n\n<br><b> Feedback-Api</b> \n <br><i>feedbackapiadm@gmail.com</i>\n<br> " + time.Now().Format(timeFormat)

	//Email send function
	if !services.SendEmail(user.Email, "Feedback request.", bodyString) {
		c.String(http.StatusBadRequest, "An error has ocurred sending the email")
		return
	}

	c.String(http.StatusCreated, "Success")
}
