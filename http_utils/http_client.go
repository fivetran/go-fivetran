package httputils

import "net/http"

// HttpClient performs an HTTP request. Can be implemented by mocks
// to perform convenient unit tests
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
