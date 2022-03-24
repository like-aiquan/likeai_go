package error

type ResponseError struct {
	Err  error
	Code int
}
