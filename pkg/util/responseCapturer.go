package util

import (
	"net/http"
)

type ResponseCapturer struct {
	http.ResponseWriter
	StatusCode  int
	Body        []byte
	WroteHeader bool
}

func NewResponseCapturer(w http.ResponseWriter) *ResponseCapturer {
	return &ResponseCapturer{ResponseWriter: w}
}

func (c *ResponseCapturer) WriteHeader(statusCode int) {
	if !c.WroteHeader {
		c.StatusCode = statusCode
		c.WroteHeader = true
		// Do NOT write to the underlying ResponseWriter
	}
}

func (c *ResponseCapturer) Write(b []byte) (int, error) {
	if !c.WroteHeader {
		c.WriteHeader(http.StatusOK)
	}
	c.Body = append(c.Body, b...) // Append in case Write is called multiple times
	// Do NOT write to the underlying ResponseWriter
	return len(b), nil
}
