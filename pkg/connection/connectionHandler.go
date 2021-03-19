package connection

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/DarthCucumber/gofuzz/pkg/utils"
)

//checks if host is reachable or not
func IsReachable(url string) {
	_, err := http.Head(url)
	utils.CheckErr(err, "[x] The URL `", url, "` with option -u is either not reachable or malformed. Pls Re-check")
}

// make concurrent requests
func MakeReq(url, method string, Wg *sync.WaitGroup) {
	req, err := http.NewRequest(method, url, nil)
	utils.CheckErr(err, err)
	client := http.Client{}
	res, err := client.Do(req)
	utils.CheckErr(err, err)
	defer res.Body.Close()
	outMsg := fmt.Sprintf("[%d] %s", res.StatusCode, url)
	fmt.Println(outMsg)
	Wg.Done()
}
