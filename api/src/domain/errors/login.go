package errors

import (
	"net/http"
)

// AuthorizationError erro de autorização
var AuthorizationError error = NewAPIValidationError(http.StatusBadRequest, "Invalid username or password")
