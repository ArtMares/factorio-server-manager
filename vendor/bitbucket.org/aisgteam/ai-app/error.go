package aiapp

import "fmt"

type Error string

func (e Error) Error() string {
    return fmt.Sprintf("aiapp: %v", string(e))
}

func Err(method string, msg interface{}) error {
    t := fmt.Sprintf("%T", msg)
    switch t {
    case "error":
        return Error(fmt.Sprintf("%s() %s", method, msg.(error).Error()))
    default:
        return Error(fmt.Sprintf("%s() %v", method, msg))
    }
}