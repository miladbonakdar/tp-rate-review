package fail

import "fmt"

type Fail struct {
	Message string `json:"message"` // Human readable message for clients
	Code    int    `json:"-"`       // HTTP Status code. We use `-` to skip json marshaling.
	Err     error  `json:"-"`       // The original error. Same reason as above.
}

func (ex Fail) Error() string {
	if ex.Err != nil {
		if ex.Message == "" {
			return fmt.Sprintf("%s \t -> %s", ex.Message, ex.Err.Error())
		} else {
			return ex.Err.Error()
		}
	}
	return ex.Message
}

func (ex Fail) Unwrap() error {
	return ex.Err
}

func (ex *Fail) WithError(err error) *Fail {
	ex.Err = err
	return ex
}

func NewFailByError(code int, err error, message string) error {
	return Fail{
		Message: message,
		Code:    code,
		Err:     err,
	}
}

func NewFail(code int, message string) error {
	return Fail{
		Message: message,
		Code:    code,
		Err:     nil,
	}
}
