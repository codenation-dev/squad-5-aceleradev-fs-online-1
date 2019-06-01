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
