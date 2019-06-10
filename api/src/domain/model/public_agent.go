package model

import "time"

// PublicAgent struct
type PublicAgent struct {
	ID         string    `json:"id,omitempty" xorm:"varchar(26) pk 'id'"`
	Name       string    `json:"name,omitempty" xorm:"varchar(100) notnull unique(identifier)"`
	Occupation string    `json:"occupation,omitempty" xorm:"varchar(160) notnull unique(identifier)"`
	Department string    `json:"department,omitempty" xorm:"varchar(160) notnull unique(identifier)"`
	Salary     float64   `json:"salary,omitempty"`
	Checked    time.Time `json:"-" xorm:"notnull"`
	CreatedAt  time.Time `json:"-" xorm:"notnull"`
	UpdatedAt  time.Time `json:"-" xorm:"notnull"`
}
