package helper

import "fmt"

type CustomError struct {
	Status  int
	Message string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("%d: %s", c.Status, c.Message)
}

func NewCustomError(status int, message string) error {
	return &CustomError{
		Status:  status,
		Message: message,
	}
}
