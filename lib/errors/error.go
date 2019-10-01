package errors

import "fmt"

type Error struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
}

func (cErr *Error) Error() string {
	return fmt.Sprintf(`{"status":%v ,"reason":"%s"}`, cErr.Status, cErr.Reason)
}
