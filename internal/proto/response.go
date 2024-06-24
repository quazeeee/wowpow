package proto

type Response struct {
	Challenge *Challenge
	Error     *Error
	Result    *Quote
}
