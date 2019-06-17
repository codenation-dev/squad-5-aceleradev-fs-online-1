package validator

// IDURI struct
type IDURI struct {
	ID string `uri:"id" binding:"required,len=26"`
}
