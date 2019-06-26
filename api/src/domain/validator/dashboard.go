package validator

import "time"

// DashboardRequest struct
type DashboardRequest struct {
	MonthStart time.Time `form:"month_start" binding:"omitempty" time_format:"2006-01"`
	MonthEnd   time.Time `form:"month_end" binding:"omitempty" time_format:"2006-01"`
}

// DashboardCustomerRequest struct
type DashboardCustomerRequest struct {
	Limit  int `form:"limit" binding:"omitempty,max=50"`
	Offset int `form:"offset" binding:"omitempty"`
}
