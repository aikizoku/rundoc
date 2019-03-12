package repository

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type httpClient struct {
	timeout time.Duration
}

func (r *httpClient) Get(url string, params map[string]interface{}, headers map[string]string) (int64, int, []byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value.(string))
	}
	req.URL.RawQuery = query.Encode()
	return r.send(req)
}

func (r *httpClient) Post(url string, params []byte, headers map[string]string) (int64, int, []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		panic(err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return r.send(req)
}

func (r *httpClient) Put(url string, params []byte, headers map[string]string) (int64, int, []byte) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(params))
	if err != nil {
		panic(err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return r.send(req)
}

func (r *httpClient) Delete(url string, params map[string]interface{}, headers map[string]string) (int64, int, []byte) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value.(string))
	}
	req.URL.RawQuery = query.Encode()
	return r.send(req)
}

func (r *httpClient) send(req *http.Request) (int64, int, []byte) {
	client := http.Client{}
	client.Timeout = r.timeout

	beforeTime := r.timeNowUnix()
	res, err := client.Do(req)
	afterTime := r.timeNowUnix()
	if err != nil {
		panic(err)
	}
	resTime := afterTime - beforeTime

	if res.StatusCode != http.StatusOK {
		return resTime, res.StatusCode, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return resTime, res.StatusCode, body
}

func (r *httpClient) timeNowUnix() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// NewHTTPClient ... リポジトリを作成する
func NewHTTPClient() HTTPClient {
	timeout := 20 * time.Second
	return &httpClient{
		timeout: timeout,
	}
}
