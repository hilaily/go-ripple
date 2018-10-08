package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var (
	client = &http.Client{}
)

func HttpGet(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	return res, err
}

func HttpPost(url string, dataJson []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(dataJson))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}
