package keycloak

// APIError holds message and statusCode for api errors
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error stringifies the APIError
func (e APIError) Error() string {
	return e.Message
}

// Get error code
func (e APIError) GetCode() int {
	return e.Code
}
