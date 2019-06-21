package model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// AlertType Enum
type AlertType int

const (
	// PublicAgentType Enum Value
	PublicAgentType AlertType = iota + 1
	// BiggerSalaryType Enum Value
	BiggerSalaryType
	// BankEmployeeType Enum Value
	BankEmployeeType
)

var toString = map[AlertType]string{
	PublicAgentType:  "PUBLIC_AGENT",
	BiggerSalaryType: "BIGGER_SALARY",
	BankEmployeeType: "BANK_EMPLOYEE",
}

var toID = map[string]AlertType{
	"PUBLIC_AGENT":  PublicAgentType,
	"BIGGER_SALARY": BiggerSalaryType,
	"BANK_EMPLOYEE": BankEmployeeType,
}

// String converte para texto
func (a AlertType) String() string {
	return toString[a]
}

// MarshalJSON marshals the enum as a quoted json string
func (a AlertType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(a.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (a *AlertType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*a = toID[j]
	return nil
}

// AlertTypeFromString converte string para AlertType
func AlertTypeFromString(v string) (*AlertType, error) {
	a := AlertType(-1)
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(v)
	buffer.WriteString(`"`)
	err := a.UnmarshalJSON(buffer.Bytes())
	if err != nil {
		return nil, err
	}
	if a <= 0 {
		return nil, fmt.Errorf(`Invalid AlertType "%s"`, v)
	}
	return &a, nil
}
