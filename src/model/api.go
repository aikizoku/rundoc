package model

// API ...
type API struct {
	Name        string
	Description string
	Endpoints   *APIEndpoints
	Request     *APIRequest
	Response    *APIResponse
}

// APIEndpoints ...
type APIEndpoints struct {
	Local      string
	Staging    string
	Production string
}

// APIRequest ...
type APIRequest struct {
	Method  string
	Path    string
	Headers string
	Params  string
}

// APIResponse ...
type APIResponse struct {
	Time       int64
	StatusCode int
	Body       string
}
