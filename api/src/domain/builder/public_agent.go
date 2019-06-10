package builder

import (
	"app/domain/model"
	"strconv"
	"time"
)

// PublicAgentFrom converte a linha para funcionário publico
func PublicAgentFrom(line *[]string) *model.PublicAgent {
	now := time.Now()
	s, _ := strconv.ParseFloat((*line)[3], 32)
	return &model.PublicAgent{
		Name:       (*line)[0],
		Occupation: (*line)[1],
		Department: (*line)[2],
		Salary:     s,
		Checked:    now,
		UpdatedAt:  now,
		CreatedAt:  now,
	}
}
