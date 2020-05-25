package throttle

import (
	"io"
	"net/http"
	"sync"
	"time"
)

type simple struct {
	lastRequest time.Time
	duration    time.Duration
	m           sync.Mutex
}

func NewSimpleThrottler(duration time.Duration) Throttler {
	return &simple{
		lastRequest: time.Now(),
		m:           sync.Mutex{},
		duration:    duration,
	}
}

// Get returns the io.ReadCloser from the resp. It is up to the consumer to properly close the ReadCloser
func (t *simple) Get(url string) (io.ReadCloser, error) {
	//start := time.Now()
	t.m.Lock()
	defer t.m.Unlock()
	if time.Since(t.lastRequest) < t.duration {
		time.Sleep(t.duration - time.Since(t.lastRequest))
	}
	resp, err := http.Get(url)
	t.lastRequest = time.Now()
	//fmt.Println(time.Since(start), url)
	return resp.Body, err
}
