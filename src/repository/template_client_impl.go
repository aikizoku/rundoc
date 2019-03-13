package repository

import (
	"bytes"

	"github.com/alecthomas/template"
)

type templateClient struct {
}

// GetMarged ... 任意の値をマージした文字列を返す
func (r *templateClient) GetMarged(tmpl string, src interface{}) string {
	t, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer
	if err := t.Execute(&doc, src); err != nil {
		panic(err)
	}
	return doc.String()
}

// NewTemplateClient ... リポジトリを作成する
func NewTemplateClient() TemplateClient {
	return &templateClient{}
}
