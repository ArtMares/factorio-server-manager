package aiapic

import (
    "io"
    "net/http"
    "net/url"
)

type ClientInterface interface {
    New(string) error
    FullName() string
    PrepareUri(string) *url.URL
    Prepare(string, *url.URL, io.Reader) (*http.Request, error)
    Exec(*http.Request) (*http.Response, error)
    SetAuth(AuthInterface)
    SetHeaders(*http.Request, map[string]string)
    SetHeader(*http.Request, string, string)
    Err(string, string) error
}

type AuthInterface interface {
    Use(*http.Request)
}