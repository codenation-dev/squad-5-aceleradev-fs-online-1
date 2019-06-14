package model

import "html/template"

// Email struct
type Email struct {
	Recipients []string
	Subject    string
	Body       template.Template
}
