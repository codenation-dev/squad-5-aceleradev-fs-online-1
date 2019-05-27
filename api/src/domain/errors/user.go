package errors

import (
	"net/http"
)

// DuplicatedUserError erro quando usuário já existe
var DuplicatedUserError error = NewAPIValidationError(http.StatusBadRequest, "Usuário já existe")
