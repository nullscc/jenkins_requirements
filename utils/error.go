package utils

type StringError struct {
	Msg string
}

func (e *StringError) Error() string {
	return e.Msg
}
