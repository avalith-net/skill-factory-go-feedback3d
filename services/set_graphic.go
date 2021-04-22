package services

import (
	"fmt"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"github.com/fatih/structs"
)

func InitGraphic(fb models.Feedback, user *models.ReturnUser) error {

	graphicAreaStats := []*[]models.MetricsCount{&user.Graphic.TechStats, &user.Graphic.TeamStats, &user.Graphic.PerfoStats}
	areas := []*structs.Struct{structs.New(fb.TechArea), structs.New(fb.TeamArea), structs.New(fb.PerformanceArea)}

	for i, area := range areas {

		graphicAreaStat := *graphicAreaStats[i]
		var container []models.MetricsCount

		for _, fieldName := range area.Names() {
			field := area.Field(fieldName)

			var found bool
			for j, c := range *graphicAreaStats[i] {
				if c.Metric == field.Value().(string) {
					graphicAreaStat[j].Count++
					found = true
					break
				}
			}

			if !found {
				found = false
				for x, v := range container {
					if v.Metric == field.Value().(string) {
						container[x].Count++
						found = true
						fmt.Println(container, v)
						break
					}
				}
				if !found {
					mc := models.MetricsCount{Metric: field.Value().(string), Count: 1}
					container = append(container, mc)
				}
			}

		}
		if len(container) != 0 {
			for _, v := range container {
				*graphicAreaStats[i] = append(*graphicAreaStats[i], v)
			}
		}
	}

	return nil
}
