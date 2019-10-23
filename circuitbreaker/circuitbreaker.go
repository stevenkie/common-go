package circuitbreaker

import (
	"time"

	"github.com/eapache/go-resiliency/breaker"
)

// ErrBreakerOpen used for breaker open error
var ErrBreakerOpen = breaker.ErrBreakerOpen

// CircuitBreaker struct for circuitbreaker
type CircuitBreaker struct {
	breaker *breaker.Breaker
	name    string
}

// New used to create circuit breaker
// errorThreshold is how many times error before CB open
// successThreshold is how many times success needed before CB become closed from half-open
// timeout is how long CB will open before change to half-open (in millisecond)
func New(errorThreshold int, successThreshold int, timeout int) CircuitBreaker {
	cb := breaker.New(errorThreshold, successThreshold, time.Duration(timeout)*time.Millisecond)

	return CircuitBreaker{
		breaker: cb,
	}
}

// Run used for running function with circuit breaker
// This will return error either circuit breaker open or the error from your function
// You can check if circuit breaker open from ErrBreakerOpen
func (cb CircuitBreaker) Run(f func() error) error {
	result := cb.breaker.Run(f)

	switch result {
	case nil:
		// Success
	case breaker.ErrBreakerOpen:
		// our function wasn't run because the breaker was open
	default:
		// other error
	}

	return result
}
