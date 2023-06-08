package service

import "github.com/aikizoku/rundoc/src/model"

type Runner interface {
	ShowRunList() error

	GetRunList() ([]string, error)

	GetRunPreview(
		name string,
	) (string, error)

	Run(
		name string,
		env string,
		doc bool,
	) (*model.API, error)
}
