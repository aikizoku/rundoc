package repository

import (
	"bytes"
	"html/template"

	"github.com/aikizoku/rundoc/src/log"
)

type templateClient struct {
}

func NewTemplateClient() TemplateClient {
	return &templateClient{}
}

func (r *templateClient) GetMerged(
	tmpl string,
	src any,
) (string, error) {
	funcMap := template.FuncMap{
		// HTMLエスケープ
		"safe_html": func(text string) template.HTML {
			return template.HTML(text)
		},
	}
	t, err := template.New("tmpl").Funcs(funcMap).Parse(tmpl)
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
