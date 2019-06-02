package errors

import (
	"net/http"
)

// DuplicatedCustomerError erro quando customer já existe
var DuplicatedCustomerError error = NewAPIValidationError(http.StatusBadRequest, "Customer já existe")

// AllDuplicatedCustomerError erro quando todos os customers da lista já exitem
var AllDuplicatedCustomerError error = NewAPIValidationError(http.StatusBadRequest, "Todos os Customer da lista já existem")

// ListDuplicatedCustomerError erro quando alguns customers da lista já exitem
var ListDuplicatedCustomerError error = NewAPIValidationError(http.StatusPartialContent, "Alguns os Customer da lista já existem")
