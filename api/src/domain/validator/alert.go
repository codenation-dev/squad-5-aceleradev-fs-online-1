package validator

import (
	"app/domain/model"
	"time"
)

// IDURI struct
type IDURI struct {
	ID string `uri:"id" binding:"required,len=26"`
}

// AgentListRequest struct
type AlertListRequest struct {
	Name       string          `form:"name" binding:"omitempty,max=100"`
	Email      string          `form:"email" binding:"omitempty,max=160"`
	Customer   string          `form:"customer" binding:"omitempty,max=100"`
	Type       model.AlertType `form:"type" binding:"omitempty,max=100"`
	DateStart  time.Time       `form:"date_start" binding:"omitempty"`
	DateFinish time.Time       `form:"date_finish" binding:"omitempty"`
	Limit      int             `form:"limit" binding:"omitempty,max=50"`
	Offset     int             `form:"offset" binding:"omitempty"`
}
