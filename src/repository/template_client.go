package repository

type TemplateClient interface {
	GetMerged(
		tmpl string,
		src any,
	) (string, error)
}
