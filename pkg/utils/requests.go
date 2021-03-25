package utils

import (
	"net/http"
	"sync"
	"time"
)

// make concurrent requests
func MakeRequest(method, url string, t int, out chan int, Wg *sync.WaitGroup) {
	defer Wg.Done()
	// TODO: add client timeout
	req, err := http.NewRequest(method, url, nil)
	CheckErr(err, err)
	//setting timeout
	client := http.Client{
		Timeout: time.Duration(t) * time.Millisecond,
	}
	res, err := client.Do(req)
	CheckErr(err, err)
	defer res.Body.Close()
	out <- res.StatusCode
}
