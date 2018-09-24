package zaipher

type APIError struct {
	Msg string
}

func (err *APIError) Error() string {
	return err.Msg
}

func newAPIError(m string) *APIError {
	return &APIError{Msg: m}
}
