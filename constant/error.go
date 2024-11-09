package constant

const (
	ErrUniqueName = "product name already registered"

	// gRPC Input Validation
	ErrValidationEmptyName        string = "product name can't be empty"
	ErrValidationEmptyDescription string = "product description can't be empty"

	// gRPC Interceptor
	ErrValidationEmptyAuth    string = "missing authentication"
	ErrValidationInvalidToken string = "invalid token"
)
