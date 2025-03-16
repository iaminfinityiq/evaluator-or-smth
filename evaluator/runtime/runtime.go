package runtime

import "fmt"

type Error interface {
	ErrorType() string
	Reason_() string
}

func DisplayError(e Error) {
	fmt.Println(e.ErrorType() + ": " + e.Reason_())
}

type RuntimeResult struct {
	Result any
	Error *Error
}

func Success(result any) RuntimeResult {
	return RuntimeResult{
		result,
		nil,
	}
}

func Failure(_error Error) RuntimeResult {
	return RuntimeResult{
		nil,
		&_error,
	}
}