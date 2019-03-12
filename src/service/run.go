package service

import "github.com/aikizoku/rundoc/src/model"

// Run ...
type Run interface {
	ShowList()
	Run(name string, env string) *model.API
}
