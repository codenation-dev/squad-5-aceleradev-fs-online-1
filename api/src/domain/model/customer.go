package model

import (
	"time"
)

// Customer struct
type Customer struct {
	ID        string    `json:"-" xorm:"varchar(26) pk 'id'"`
	Name      string    `json:"-" xorm:"varchar(25) notnull unique"`
	Salary    float32   `json:"-" xorm:"descimal"`
	CreatedAt time.Time `json:"-" xorm:"notnull created"`
	UpdatedAt time.Time `json:"-" xorm:"notnull updated"`
}

// CustomerList struct
type CustomerList struct {
	Records int64
	Data    []Customer
}
