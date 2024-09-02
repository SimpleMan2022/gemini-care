package errorHandler

type BadRequestError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (err *InternalServerError) Error() string {
	return err.Message
}

func (err *BadRequestError) Error() string {
	return err.Message
}

func (err *NotFoundError) Error() string {
	return err.Message
}
