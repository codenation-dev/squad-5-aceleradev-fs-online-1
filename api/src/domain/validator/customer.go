package validator

// CustomerURI struct
type CustomerURI struct {
	CustomerID string `uri:"customerId" binding:"required,len=26"`
}

// CustomerCreation struct
type CustomerCreation struct {
	Name   string  `json:"name" binding:"required,max=50,min=1"`
	Salary float32 `json:"salary" binding:"required,min=1"`
}

// CustomerListRequest struct
type CustomerListRequest struct {
	Limit  int `form:"limit" binding:"omitempty,max=50"`
	Offset int `form:"offset" binding:"omitempty"`
}
