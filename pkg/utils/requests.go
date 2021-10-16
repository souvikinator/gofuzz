package utils

import (
	"math"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// make concurrent requests
func MakeRequest(method, url string, customTransport *http.Transport, t, rt int, out chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		res     *http.Response
		err     error
		retries = rt + 1
		client  = http.Client{
			Transport: customTransport,
			Timeout:   time.Duration(t) * time.Millisecond,
		}
	)

	req, err := http.NewRequest(method, url, nil)
	CheckErr(err, err)

	//try to execute the request until it succeeds or extinguishes all retries
	for r := 0; r < retries; r++ {
		if res, err = client.Do(req); err == nil {
			break
		}
		backoff := int64(math.Pow(2, float64(r))) * 10
		time.Sleep(time.Duration(backoff) * time.Millisecond)
	}

	//check for timeout error and return
	if e, ok := err.(net.Error); ok && e.Timeout() {
		out <- []string{"timeout", url}
		return
	} else if err != nil {
		ShowError(err, err)
		os.Exit(0)
	}
	defer res.Body.Close()
	out <- []string{strconv.Itoa(res.StatusCode), url}
}
