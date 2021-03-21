package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

// make concurrent requests
//TODO: add timeout feature
func Fuzz(parsedUrl []string, fuzzdata, method string, out chan []string, Wg *sync.WaitGroup) {
	defer Wg.Done()
	//Request part
	newUrl := strings.Join(parsedUrl, url.PathEscape(fuzzdata))
	//unnecessary but can't find an alternative :(
	outUrl := strings.Join(parsedUrl, fuzzdata)
	req, err := http.NewRequest(method, newUrl, nil)
	CheckErr(err, err)
	client := http.Client{}
	res, err := client.Do(req)
	CheckErr(err, err)
	defer res.Body.Close()
	out <- []string{strconv.Itoa(res.StatusCode), fuzzdata}
	outMsg := fmt.Sprintf("[%d] %s", res.StatusCode, outUrl)
	fmt.Println(outMsg)
}
