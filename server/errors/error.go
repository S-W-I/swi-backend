package errors

type AppError string
const (
	UnsupportedMethodError AppError = "method not supported"
)

func (err AppError) Error() string {
	return string(err)
}