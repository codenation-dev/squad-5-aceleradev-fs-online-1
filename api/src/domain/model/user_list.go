package model

// UserList struct
type UserList struct {
	Records int64  `json:"records,omitempty"`
	Data    []User `json:"data,omitempty"`
}
