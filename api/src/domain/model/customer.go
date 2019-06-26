package model

import (
	"time"
)

// Customer struct
type Customer struct {
	ID        string    `json:"id,omitempty" xorm:"varchar(26) pk 'id'"`
	Name      string    `json:"name" xorm:"varchar(50) notnull unique"`
	Salary    float32   `json:"salary" xorm:"salary"`
	CreatedAt time.Time `json:"-" xorm:"notnull created"`
	UpdatedAt time.Time `json:"-" xorm:"notnull updated"`
}

// CustomerList struct
type CustomerList struct {
	Records int64      `json:"records"`
	Data    []Customer `json:"data"`
}

// CustomerInsert struct
type CustomerInsert struct {
	Success      int `json:"success"`
	AlreadyExist int `json:"alreadyExist"`
}
