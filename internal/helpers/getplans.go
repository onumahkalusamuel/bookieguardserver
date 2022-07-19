package helpers

import (
	"strconv"

	"bookieguardserver/internal/models"
)

type Plan struct {
	Key      string
	Title    string
	Duration uint
	Price    uint
}

func GetPlans() []Plan {
	var returnPlans []Plan
	plans := []Plan{
		{Key: "plan1", Title: "3 Months", Duration: 3},
		{Key: "plan2", Title: "6 Months", Duration: 6},
		{Key: "plan3", Title: "1 Year", Duration: 12},
		{Key: "plan4", Title: "2 Years", Duration: 24},
		{Key: "plan5", Title: "Lifetime", Duration: 240},
	}

	// get plans
	for _, plan := range plans {
		s := models.Settings{}
		s.Setting = plan.Key
		s.Read()

		converted, _ := strconv.Atoi(s.Value)

		plan.Price = uint(converted)

		returnPlans = append(returnPlans, plan)
	}

	return returnPlans
}
