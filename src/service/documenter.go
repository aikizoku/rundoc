package service

import "github.com/aikizoku/rundoc/src/model"

type Documenter interface {
	Distribute(
		name string,
		api *model.API,
	) error
}
