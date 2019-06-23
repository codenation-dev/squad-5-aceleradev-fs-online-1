package model

// Dashboard struct
type Dashboard struct {
	Totals Totals          `json:"totals"`
	Data   []DashboardData `json:"data"`
}

// Totals struct
type Totals struct {
	CustomerQuantity    int64 `json:"customer_quantity"`
	PublicAgentQuantity int64 `json:"public_agent_quantity"`
	EmployeeQuantity    int64 `json:"employee_quantity"`
}

// PublicAgentData struct
type PublicAgentData struct {
	CustomerQuantity     int64 `json:"customer_quantity"`
	NotifyQuantity       int64 `json:"notify_quantity"`
	BiggerSalaryQuantity int64 `json:"bigger_salary_quantity"`
	NewQuantity          int64 `json:"new_quantity"`
}

// BiggerSalary struct
type BiggerSalary struct {
	CustomerQuantity int64 `json:"customer_quantity"`
	NotifyQuantity   int64 `json:"notify_quantity"`
}

// BankEmployee struct
type BankEmployee struct {
	CustomerQuantity int64 `json:"customer_quantity"`
	NotifyQuantity   int64 `json:"notify_quantity"`
	NewQuantity      int64 `json:"new_quantity"`
}

// Clients struct
type Clients struct {
	NewQuantity int64 `json:"new_quantity"`
}

// Alerts struct
type Alerts struct {
	PublicAgent  PublicAgentData `json:"public_agent"`
	BiggerSalary BiggerSalary    `json:"bigger_salary"`
	BankEmployee BankEmployee    `json:"bank_employee"`
	Clients      Clients         `json:"clients"`
}

// DashboardData struct
type DashboardData struct {
	Month  string `json:"month"`
	Alerts Alerts `json:"alerts"`
}
