package service

import "github.com/aikizoku/rundoc/src/model"

// Doc ...
type Doc interface {
	Distribute(name string, api *model.API)
}
