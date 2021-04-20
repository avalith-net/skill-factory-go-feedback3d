package services

import (
	"github.com/blotin1993/feedback-api/models"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func InitGraphic(fb models.Feedback, user models.ReturnUser) (models.Graphic, error) {

	graphicAreaMaps := []map[string]interface{}{structs.Map(user.Graphic.Tech), structs.Map(user.Graphic.Team), structs.Map(user.Graphic.Perfo)}
	graphicPointers := []interface{}{&user.Graphic.Tech, &user.Graphic.Team, &user.Graphic.Perfo}
	areas := []*structs.Struct{structs.New(fb.TechArea), structs.New(fb.TeamArea), structs.New(fb.PerformanceArea)}

	for i, area := range areas {

		graphicAreaMap := graphicAreaMaps[i]

		for _, fieldName := range area.Names() {
			field := area.Field(fieldName)

			var num int
			//-------------------------------------
			switch field.Value().(string) {
			case "Let´s Work On This":
				num = 25
			case "Reach The Goal":
				num = 50
			case "Relevant Performance":
				num = 75
			case "Master":
				num = 100
			default:
				num = 0
			}
			var custom models.Custom
			err := mapstructure.Decode(graphicAreaMap[fieldName], &custom)
			if err != nil {
				return models.Graphic{}, err
			}
			if num == 0 {
				custom.Count += 0
			} else {
				custom.Count++
			}
			custom.Value += num
			graphicAreaMap[fieldName] = custom

		}
		err := mapstructure.Decode(graphicAreaMap, graphicPointers[i])
		if err != nil {
			return models.Graphic{}, err
		}

	}

	return user.Graphic, nil
}
