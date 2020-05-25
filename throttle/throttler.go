package throttle

import (
	"io"
	"time"
)

const (
	DefaultThrottleDuration = time.Millisecond * 200
)

type Throttler interface {
	Get(url string) (io.ReadCloser, error)
}
