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

func NewErrorUseCaseResponse(errorName string, errorMessage string) UseCaseResponse {
	return NewUseCaseResponse(nil, &errorName, &errorMessage)
}

func NewSuccessUseCaseResponse(data any) UseCaseResponse {
	return NewUseCaseResponse(data, nil, nil)
}
