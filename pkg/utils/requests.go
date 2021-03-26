package utils

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

// make concurrent requests
func MakeRequest(method, url string, t int, out chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
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
	out <- []string{strconv.Itoa(res.StatusCode), url}
}
