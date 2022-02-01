package repository

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aikizoku/rundoc/src/log"
)

type httpClient struct {
	timeout time.Duration
}

func NewHTTPClient() HTTPClient {
	return &httpClient{
		10 * time.Minute,
	}
}

func (r *httpClient) Get(url string, params map[string]interface{}, headers map[string]string) (int64, int, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf(err, "HTTP Request作成に失敗: %s", url)
		return 0, 0, nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	query := req.URL.Query()
	for key, value := range params {
		if v, ok := value.(string); ok {
			query.Add(key, v)
		}
		if v, ok := value.(int64); ok {
			query.Add(key, fmt.Sprintf("%d", v))
		}
		if v, ok := value.(float64); ok {
			query.Add(key, fmt.Sprintf("%f", v))
		}
		if v, ok := value.(bool); ok {
			query.Add(key, fmt.Sprintf("%t", v))
		}
	}
	req.URL.RawQuery = query.Encode()
	return r.send(req)
}

func (r *httpClient) Post(url string, params []byte, headers map[string]string) (int64, int, []byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		log.Errorf(err, "HTTP Request作成に失敗: %s", url)
		return 0, 0, nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return r.send(req)
}

func (r *httpClient) Put(url string, params []byte, headers map[string]string) (int64, int, []byte, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(params))
	if err != nil {
		log.Errorf(err, "HTTP Request作成に失敗: %s", url)
		return 0, 0, nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return r.send(req)
}

func (r *httpClient) Delete(url string, params map[string]interface{}, headers map[string]string) (int64, int, []byte, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Errorf(err, "HTTP Request作成に失敗: %s", url)
		return 0, 0, nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	query := req.URL.Query()
	for key, value := range params {
		if v, ok := value.(string); ok {
			query.Add(key, v)
		}
		if v, ok := value.(int64); ok {
			query.Add(key, fmt.Sprintf("%d", v))
		}
		if v, ok := value.(float64); ok {
			query.Add(key, fmt.Sprintf("%f", v))
		}
		if v, ok := value.(bool); ok {
			query.Add(key, fmt.Sprintf("%t", v))
		}
	}
	req.URL.RawQuery = query.Encode()
	return r.send(req)
}

func (r *httpClient) send(req *http.Request) (int64, int, []byte, error) {
	client := http.Client{}
	client.Timeout = r.timeout

	beforeTime := r.timeNowUnix()
	res, err := client.Do(req)
	afterTime := r.timeNowUnix()
	if err != nil {
		log.Errorf(err, "HTTP Request送信に失敗: %s", req.URL.String())
		return 0, 0, nil, err
	}
	resTime := afterTime - beforeTime

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf(err, "Response Bodyの読み込みに失敗")
		return resTime, res.StatusCode, nil, err
	}
	defer res.Body.Close()

	return resTime, res.StatusCode, body, nil
}

func (r *httpClient) timeNowUnix() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
