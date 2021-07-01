package repository

import (
	"github.com/cenkalti/backoff"
	"time"
)

func config() *backoff.ExponentialBackOff {
	back := backoff.NewExponentialBackOff()
	back.MaxInterval = 20 * time.Second
	back.Multiplier = 2
	return back
}
