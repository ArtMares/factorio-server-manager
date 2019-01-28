package aiapp

import "net/http"

type StoreInterface interface {
    Save(w http.ResponseWriter, r *http.Request) *http.Request
}