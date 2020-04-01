package network

import (
	"github.com/iochen/jandanTreeholeRSS/jandan"
	"io"
	"net/http"
)

func HttpGetWithUA(url string) (io.Reader, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", jandan.UA)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
