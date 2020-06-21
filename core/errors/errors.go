package errors

type Operation string
type ErrorType string
type ErrorLevel string

const (
	FailedToCreate ErrorType = "FAILED_TO_CREATE"
	FailedToLoad   ErrorType = "FAILED_TO_LOAD"
	Unexpected     ErrorType = "UNEXPECTED"

	StandardErrorLevel ErrorLevel = "ERROR"
)

type Error struct {
	operations []Operation
	errorType  ErrorType
	error      error
	severity   ErrorLevel
}

func NewError(operation Operation, errorType ErrorType, err error, severity ErrorLevel) *Error {
	return &Error{
		operations: []Operation{operation},
		errorType:  errorType,
		error:      err,
		severity:   severity,
	}
}

func (e *Error) AddOperation(operation Operation) *Error {
	e.operations = append(e.operations, operation)
	return e
}

func (e *Error) Operations() []Operation {
	return e.operations
}

func (e *Error) ErrorType() ErrorType {
	return e.errorType
}

func (e *Error) Error() error {
	return e.error
}

func (e *Error) Severity() ErrorLevel {
	return e.severity
}
