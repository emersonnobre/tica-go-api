package types

type UseCaseResponse struct {
	Data         any
	ErrorName    *string
	ErrorMessage *string
}

func NewUseCaseResponse(data any, errorName *string, errorMessage *string) UseCaseResponse {
	return UseCaseResponse{
		Data:         data,
		ErrorName:    errorName,
		ErrorMessage: errorMessage,
	}
}
