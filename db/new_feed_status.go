package db

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/models"
// 	"github.com/jinzhu/copier"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func AddFeedStatus(fbs models.FeedbackStatus) (string, bool, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	db := MongoCN.Database("feedback-db")
// 	col := db.Collection("feedbacks-status")

// 	var (
// 		feedStatus   models.FeedbackStatus
// 		fbRequested  models.FeedbacksRequested
// 		userAsksFeed models.UsersAskedFeed
// 		err          error
// 	)

// 	fbrIDString, _, stringErr := AddFeedbackRequested(fbRequested)
// 	if stringErr != nil {
// 		fmt.Println("Error trying to create a new feedbackRequestedStatus from feedStatus register")
// 		return "", false, err
// 	}

// 	fmt.Println("1.Imprimiendo fbrIDString")
// 	fmt.Println(fbrIDString)

// 	fbrObteined, obteinedErr := GetFeedBackRequested(fbrIDString)
// 	if obteinedErr != nil {
// 		fmt.Println("Error trying to get a feedbackRequestedStatus from feedStatus register")
// 		return "", false, err
// 	}

// 	fmt.Println("2.Imprimiendo fbrObteined")
// 	fmt.Println(fbrObteined)

// 	askedIDString, _, askedErr := AddUsersAsksFeed(userAsksFeed)
// 	if askedErr != nil {
// 		fmt.Println("Error trying to create a new userAsksfeed from feedStatus register")
// 		return "", false, err
// 	}

// 	fmt.Println("3.Imprimiendo askedIDString")
// 	fmt.Println(askedIDString)

// 	askedObteined, obteinedErr := GetUsersAskedFeedback(askedIDString)
// 	if obteinedErr != nil {
// 		fmt.Println("Error trying to get a feedbackRequestedStatus from feedStatus register")
// 		return "", false, err
// 	}

// 	fmt.Println("4.Imprimiendo askedObteined")
// 	fmt.Println(askedObteined)

// 	copier.Copy(&feedStatus, &fbs)

// 	feedStatus.FeedbacksRequested = fbrObteined
// 	feedStatus.UsersAskedFeed = askedObteined

// 	// Creamos 2 struct, uno de tipo FeedbackRequested y otro de tipo UsersAsksFeedback.
// 	//los inserto en la bd en sus respectivas colecciones, les hago un get para que
// 	//me devuelvan el obj completo(no el ID en string).
// 	// Hago una copia con el copier de la estructura. Una vez hecho eso, los agrego
// 	//(sino funciona asignandolo como ahora, probar con modify)
// 	// a ambos objetos a estructura vacia feedStatus. Y por ultimo inserto
// 	//feedStatus en la bd, no fbs, ya que viene incompleto.

// 	result, err := col.InsertOne(ctx, feedStatus)
// 	if err != nil {
// 		return "", false, err
// 	}
// 	ObjID, _ := result.InsertedID.(primitive.ObjectID)
// 	return ObjID.Hex(), true, nil
// }
