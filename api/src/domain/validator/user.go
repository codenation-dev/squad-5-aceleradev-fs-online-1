package validator

// UserURI struct
type UserURI struct {
	UserID string `uri:"userId" binding:"required,len=26"`
}

// UserCreation struct
type UserCreation struct {
	Username string `json:"username" binding:"required,max=25,min=1"`
	Password string `json:"password" binding:"required,max=20,min=6"`
	Name     string `json:"name" binding:"required,max=100,min=1"`
	Email    string `json:"email" binding:"required,max=160,min=3"`
}

// UserListRequest struct
type UserListRequest struct {
	Name   string `form:"name" binding:"omitempty,max=100"`
	Email  string `form:"email" binding:"omitempty,max=160"`
	Limit  int    `form:"limit" binding:"omitempty,max=50"`
	Offset int    `form:"offset" binding:"omitempty"`
}
