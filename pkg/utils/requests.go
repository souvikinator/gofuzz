package utils

import (
	"net"
	"net/http"
	"os"
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
