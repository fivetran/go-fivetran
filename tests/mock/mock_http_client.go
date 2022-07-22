// Copyright 2022 Fivetran

package mock

// *** Mock HTTP client for Go ***
//
// How to use?
// First, you need to create the client:
//
//   mockClient := mock.NewHttpClient()
//
// Then you need to stab HTTP methods/URLs you need. Stabbing can be exact:
//
//   mockClient.When(http.MethodGet, "/index.html")
//
// or you can stub all HTTP methods at once:
//
//   mockClient.WhenAnyMethod("/index.html")
//
// Alternatively you can stub using the wildcard symbols ('*' and '?'):
//
//   mockClient.WhenWc(http.MethodGet, "/v1/connectors/*")
//   mockClient.WhenAnyMethodWc("/v?/connectors/*")
//
// Now you need to return something from the stabbed methods.
// If you need just the HTTP status code, the following is enough:
//
//   mockClient.When(http.MethodGet, "/index.html").ThenReply(200)
//
// Now if you want to add the response body, you can Go with:
//
//   mockClient.When(http.MethodGet, "/index.html").ThenReply(200).WithBody("{}")
//
// Or in the Master Yoda-style:
//
//   mockClient.When(http.MethodGet, "/index.html").WithBody("{}").ThenReply(200)
//
// Alternatively you can use a supplied function to create the response:
//
//   mockClient.WhenAnyMethod("/v1/connectors/*").ThenCall(
//       func(req *http.Request) (*http.Response, error) {
//           resp := mock.NewResponse(req, "OK", 200, "{}")
//           return resp, nil
//       })
//
// Note that you can use the mock.NewResponse function to create the initial response object.
//
// Finally you can check how many times a stubbed method was called:
//
//   handler := mockClient.When(http.MethodGet, "/index.html").ThenReply(200)
//   ...
//   execute some methods being tested
//   ...
//   <Assert_that>(handler.Interactions == 1)
//
// Now you are ready to Go testing!
//
// ***

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HandlerFunc func(req *http.Request) (*http.Response, error)

type HttpClient struct {
	handlers   map[stubAddress]*Handler
	handlersWc map[stubAddress]*Handler
}

type Handler struct {
	statusCode   int
	statusText   string
	body         string
	function     HandlerFunc
	Interactions int
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		handlers:   map[stubAddress]*Handler{},
		handlersWc: map[stubAddress]*Handler{},
	}
}

func NewResponse(req *http.Request, code int, body string) *http.Response {
	return &http.Response{
		Status:           http.StatusText(code),
		StatusCode:       code,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           map[string][]string{},
		Body:             io.NopCloser(strings.NewReader(body)),
		ContentLength:    int64(len(body)),
		TransferEncoding: []string{},
		Close:            false,
		Uncompressed:     false,
		Trailer:          map[string][]string{},
		Request:          req,
		TLS:              nil,
	}
}

// ***

type stubAddress struct {
	method string
	url    string
}

// *** HttpClient

// Do mocks the http.Client.Do() function
func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {

	handler := c.findHandler(req)
	if handler != nil {
		return handler.handle(req)
	}

	return nil, fmt.Errorf("%s %s is not implemented", req.Method, req.URL.Path)
}

// Reset resets the entire HttpClient and deletes all handlers
func (c *HttpClient) Reset() *HttpClient {
	c.handlers = map[stubAddress]*Handler{}
	c.handlersWc = map[stubAddress]*Handler{}
	return c
}

// When creates and returns a handler which will be used for the
// exact match of the request to the specified HTTP method and URL path
func (c *HttpClient) When(method string, url string) *Handler {
	return c.addHandler(method, url, false)
}

// WhenAnyMethod creates and returns a handler which will be used for the
// exact match of the request to the specified URL path and any HTTP method
func (c *HttpClient) WhenAnyMethod(url string) *Handler {
	return c.addHandler("", url, false)
}

// WhenWc creates and returns a handler which will be used for the wildcard
// match of the request to the specified HTTP method and URL path pattern
func (c *HttpClient) WhenWc(method string, url string) *Handler {
	return c.addHandler(method, url, true)
}

// WhenAnyMethodWc creates and returns a handler which will be used for the
// wildcard match of the request to the specified URL path pattern and
// any HTTP method
func (c *HttpClient) WhenAnyMethodWc(url string) *Handler {
	return c.addHandler("", url, true)
}

func (c *HttpClient) addHandler(method string, url string, wildcard bool) *Handler {

	if method == "" {
		method = "*"
	}

	address := stubAddress{
		method: method,
		url:    url,
	}

	handler := &Handler{}

	if wildcard {
		c.handlersWc[address] = handler
	} else {
		c.handlers[address] = handler
	}

	return handler
}

func (c *HttpClient) findHandler(req *http.Request) *Handler {

	address := stubAddress{
		method: req.Method,
		url:    req.URL.Path,
	}

	// First try the exact match of the method + URL
	handler, ok := c.handlers[address]
	if ok {
		return handler
	}

	// Then try any method for the exact URL
	address.method = "*"
	handler, ok = c.handlers[address]
	if ok {
		return handler
	}

	// Finally check the wildcard handlers
	for k, v := range c.handlersWc {
		if WildcardMatch(address.url, k.url) && (req.Method == k.method || k.method == "*") {
			return v
		}
	}

	return nil
}

// *** Handler

// ThenReply sets the HTTP response status code for the handler.
// The HTTP response text is set automatically according to the
// response status code
func (h *Handler) ThenReply(statusCode int) *Handler {
	h.statusCode = statusCode
	return h
}

// WithCustomStatus allows to change the standard HTTP response text
// to a custom one.
func (h *Handler) WithCustomStatus(status string) *Handler {
	h.statusText = status
	return h
}

// WithBody adds body to the handler response. It can be called either
// before or after ThenReply()
func (h *Handler) WithBody(body string) *Handler {
	h.body = body
	return h
}

// ThenCall tells the handler to call provided function instead of
// simply returning the response. This function can implement
// some more complex logic
func (h *Handler) ThenCall(function HandlerFunc) *Handler {
	h.function = function
	return h
}

func (h *Handler) handle(req *http.Request) (*http.Response, error) {
	h.Interactions++

	if h.function != nil {
		return h.function(req)
	}

	response := NewResponse(req, h.statusCode, h.body)
	if h.statusText != "" {
		response.Status = h.statusText
	}

	return response, nil
}
