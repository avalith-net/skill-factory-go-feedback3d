package services

import (
	"fmt"

	"github.com/blotin1993/feedback-api/models"
	"github.com/fatih/structs"
)

func InitGraphic(fb models.Feedback, user models.ReturnUser) error {
	// init map
	var auxMap = map[string]int{
		"Let´s Work On This":   0,
		"Reach The Goal":       0,
		"Relevant Performance": 0,
		"Master":               0,
	}

	//create a structs.struct from fb to get its values.
	s := structs.New(fb)
	// get values from s
	for _, f := range s.Values() {
		str := fmt.Sprintf("%v", f)
		if str != "" {
			auxMap[str]++
		}
	}
	if len(user.Graphic) == 0 {
		for k, v := range auxMap {
			//init metrics count
			mc := models.MetricsCount{Metric: k, Count: v}
			user.Graphic = append(user.Graphic, mc)
		}
	} else {
		//init graph
		for i, s := range user.Graphic {

			//update map
			if val, ok := auxMap[s.Metric]; ok {
				user.Graphic[i].Count += val

			}
		}
	}
	return nil
}
