package repository

import (
	"bytes"

	"github.com/aikizoku/rundoc/src/log"
	"github.com/alecthomas/template"
)

type templateClient struct {
}

// GetMarged ... 任意の値をマージした文字列を返す
func (r *templateClient) GetMarged(tmpl string, src interface{}) (string, error) {
	t, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		log.Errorf(err, "テンプレートファイルのParseに失敗: %s", tmpl)
		return "", err
	}

	var doc bytes.Buffer
	if err := t.Execute(&doc, src); err != nil {
		log.Errorf(err, "テンプレートファイルの読み込みに失敗: %s", tmpl)
		return "", err
	}
	return doc.String(), nil
}

// NewTemplateClient ... リポジトリを作成する
func NewTemplateClient() TemplateClient {
	return &templateClient{}
}
