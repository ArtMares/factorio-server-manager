package wapp

import "fmt"

type Error string

func (e Error) Error() string {
    return fmt.Sprintf("wapp: %s", e)
}