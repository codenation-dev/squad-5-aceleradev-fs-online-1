package validator

// Login struct
type Login struct {
	Username string `json:"username" binding:"required,max=25,min=1"`
	Password string `json:"password" binding:"required,max=20,min=6"`
}
