package model

type FileCommon struct {
	Endpoints *FileEndpoints    `json:"endpoints"`
	Headers   map[string]string `json:"headers"`
}

type FileEndpoints struct {
	Local      string `json:"local"`
	Staging    string `json:"staging"`
	Production string `json:"production"`
}

type FileAuth struct {
	Local      string `json:"local"`
	Staging    string `json:"staging"`
	Production string `json:"production"`
}

type FileRun struct {
	Description string            `json:"description"`
	Path        string            `json:"path"`
	Method      string            `json:"method"`
	Headers     map[string]string `json:"headers"`
	Params      map[string]any    `json:"params"`
}
