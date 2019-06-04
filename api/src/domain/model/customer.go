package model

import (
	"time"
)

// Customer struct
type Customer struct {
	ID        string    `json:"id,omitempty" xorm:"varchar(26) pk 'id'"`
	Name      string    `json:"-" xorm:"varchar(50) notnull unique"`
	Salary    float32   `json:"-" xorm:"salary"`
	CreatedAt time.Time `json:"-" xorm:"notnull created"`
	UpdatedAt time.Time `json:"-" xorm:"notnull updated"`
}

// CustomerList struct
type CustomerList struct {
	Records int64
	Data    []Customer
}

// CustomerInsert struct
type CustomerInsert struct {
	Success      int `json:"success"`
	AlreadyExist int `json:"alreadyExist"`
}
