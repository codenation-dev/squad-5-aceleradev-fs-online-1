package builder

import (
	"app/domain/model"
	"strconv"
	"strings"
	"time"
)

// PublicAgentFrom converte a linha para funcionário publico
func PublicAgentFrom(line *[]string) *model.PublicAgent {
	now := time.Now()
	s, _ := strconv.ParseFloat(strings.Replace((*line)[3], ",", ".", 1), 32)
	return &model.PublicAgent{
		ID:         NewULID(),
		Name:       strings.Trim((*line)[0], "\x00"),
		Occupation: strings.Trim((*line)[1], "\x00"),
		Department: strings.Trim((*line)[2], "\x00"),
		Salary:     s,
		Checked:    now,
		UpdatedAt:  now,
		CreatedAt:  now,
	}
}
