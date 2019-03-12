package repository

// HTTPClient ... HTTP通信を行うリポジトリ
type HTTPClient interface {
	Get(url string, params map[string]interface{}, headers map[string]string) (int64, int, []byte)
	Post(url string, params []byte, headers map[string]string) (int64, int, []byte)
	Put(url string, params []byte, headers map[string]string) (int64, int, []byte)
	Delete(url string, params map[string]interface{}, headers map[string]string) (int64, int, []byte)
}
