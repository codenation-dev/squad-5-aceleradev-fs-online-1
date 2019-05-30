package model

import (
	"time"
)

// User struct
type User struct {
	ID        string    `json:"id,omitempty" xorm:"varchar(26) pk 'id'"`
	Username  string    `json:"username,omitempty" xorm:"varchar(25) notnull unique"`
	Password  string    `json:"-" xorm:"notnull"`
	Name      string    `json:"name,omitempty" xorm:"varchar(100) notnull"`
	Email     string    `json:"email,omitempty" xorm:"varchar(160) notnull unique"`
	CreatedAt time.Time `json:"-" xorm:"notnull created"`
	UpdatedAt time.Time `json:"-" xorm:"notnull updated"`
}

// UserList struct
type UserList struct {
	Records int64  `json:"records,omitempty"`
	Data    []User `json:"data,omitempty"`
}
