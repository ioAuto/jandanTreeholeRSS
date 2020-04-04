package network

import (
	"github.com/iochen/jandanTreeholeRSS/jandan"
	"io/ioutil"
	"net/http"
)

var NetworkRetry = 5
var networkRetry = -1

func HttpGetWithUA(url string) ([]byte, error) {
	if networkRetry == -1 {
		networkRetry = NetworkRetry
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", jandan.UA)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && networkRetry > 0 {
		networkRetry--
		return HttpGetWithUA(url)
	}

	b, err := ioutil.ReadAll(resp.Body)

	networkRetry = NetworkRetry
	return b, err
}
