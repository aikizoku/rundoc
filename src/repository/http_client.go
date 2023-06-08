package repository

type HTTPClient interface {
	Get(
		url string,
		params map[string]any,
		headers map[string]string,
	) (int64, int, []byte, error)

	Post(
		url string,
		params []byte,
		headers map[string]string,
	) (int64, int, []byte, error)

	Put(
		url string,
		params []byte,
		headers map[string]string,
	) (int64, int, []byte, error)

	Delete(
		url string,
		params map[string]any,
		headers map[string]string,
	) (int64, int, []byte, error)
}
