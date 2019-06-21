package model

import (
	"time"
)

// Alert struct
type Alert struct {
	ID            string       `json:"id,omitempty" xorm:"varchar(26) pk 'id'"`
	Type          AlertType    `json:"type" xorm:"int notnull"`
	Description   string       `json:"description" xorm:"varchar(100) notnull"`
	Customer      *Customer    `json:"customer,omitempty" xorm:"customer_id varchar(26) index null"`
	PublicAgent   *PublicAgent `json:"public_agent,omitempty" xorm:"public_agent_id varchar(26) index null"`
	User          *User        `json:"user,omitempty" xorm:"user_id varchar(26) index null"`
	UsersReceived []User       `json:"users_received,omitempty" xorm:"-"`
	CreatedAt     time.Time    `json:"datetime" xorm:"notnull created"`
	UpdatedAt     time.Time    `json:"-" xorm:"notnull updated"`
}

// AlertUser struct
type AlertUser struct {
	User  User  `json:"user" xorm:"user_id varchar(26) pk notnull"`
	Alert Alert `json:"alert" xorm:"alert_id varchar(26) pk notnull"`
}

// AlertList struct
type AlertList struct {
	Records int64       `json:"records,omitempty"`
	Data    []AlertItem `json:"data,omitempty"`
}

// AlertItem struct
type AlertItem struct {
	ID           string    `json:"id,omitempty" xorm:"varchar(26) pk 'id'"`
	Type         AlertType `json:"type" xorm:"int notnull"`
	CustomerName string    `json:"customer_name,omitempty" xorm:"customer_name varchar(26) index null"`
	CreatedAt    time.Time `json:"datetime" xorm:"notnull created"`
}
