package api

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	cb "github.com/sleey/common-go/circuitbreaker"
)

// Request request config for http call
// Please be aware that timeout is Duration type i.e 4*time.Second or 100*time.Millisecond
type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Timeout time.Duration
	CB      *cb.CircuitBreaker
}

// Do do http call
func (r *Request) Do() (buffer []byte, err error) {
	err = r.CB.Run(func() error {
		var errCb error
		buffer, errCb = r.DoRequest()

		return errCb
	})

	return
}

// DoRequest to do request
func (r *Request) DoRequest() (buffer []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), (r.Timeout))
	defer cancel()

	var request *http.Request
	var response *http.Response

	// setup http request
	request, err = http.NewRequest(r.Method, r.URL, &bytes.Buffer{})
	if err != nil {
		return
	}

	for key, value := range r.Headers {
		request.Header.Set(key, value)
	}
	request.WithContext(ctx)

	// send request
	response, err = (&http.Client{Timeout: r.Timeout}).Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// check non 200 responses
	if response.StatusCode != 200 {
		err = fmt.Errorf("Status code: %d, url: %s", response.StatusCode, r.URL)
		return
	}

	// read body response
	buffer, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	return
}
