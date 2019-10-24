package api

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// MockAPI create API mock
func MockAPI(wantedResult string) *httptest.Server {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		io.WriteString(rw, wantedResult)
	})

	return httptest.NewServer(handler)
}
