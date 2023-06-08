package model

type API struct {
	Name        string
	Description string
	Endpoints   *APIEndpoints
	Request     *APIRequest
	Response    *APIResponse
	Command     string
}

type APIEndpoints struct {
	Local      string
	Staging    string
	Production string
}

type APIRequest struct {
	Method  string
	Path    string
	Headers string
	Params  string
}

type APIResponse struct {
	Time       int64
	StatusCode int
	Body       string
}
