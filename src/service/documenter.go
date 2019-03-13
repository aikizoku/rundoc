package service

import "github.com/aikizoku/rundoc/src/model"

// Documenter ...
type Documenter interface {
	Distribute(name string, api *model.API)
}
