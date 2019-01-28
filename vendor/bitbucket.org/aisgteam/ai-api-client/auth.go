package aiapic

import "net/http"

type BasicAuth struct {
    Username    string
    Password    string
}

func (auth *BasicAuth) Use(r *http.Request) {
    r.SetBasicAuth(auth.Username, auth.Password)
}