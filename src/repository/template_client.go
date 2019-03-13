package repository

// TemplateClient ... テンプレートファイルに関するリポジトリ
type TemplateClient interface {
	GetMarged(tmpl string, src interface{}) string
}
