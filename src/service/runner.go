package service

import "github.com/aikizoku/rundoc/src/model"

// Runner ...
type Runner interface {
	ShowList()
	Run(name string, env string) *model.API
}
