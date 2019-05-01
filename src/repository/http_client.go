package repository

// HTTPClient ... HTTP通信を行うリポジトリ
type HTTPClient interface {
	Get(url string, params map[string]interface{}, headers map[string]string) (int64, int, []byte, error)
	Post(url string, params []byte, headers map[string]string) (int64, int, []byte, error)
	Put(url string, params []byte, headers map[string]string) (int64, int, []byte, error)
	Delete(url string, params map[string]interface{}, headers map[string]string) (int64, int, []byte, error)
}
