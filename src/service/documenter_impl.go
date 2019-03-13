package service

import (
	"github.com/aikizoku/rundoc/src/model"
	"github.com/aikizoku/rundoc/src/repository"
)

type documenter struct {
	configDir string
	docsDir   string
	fRepo     repository.File
	tRepo     repository.TemplateClient
}

func (s *documenter) Distribute(name string, api *model.API) {
	b := getBinFileData("doc.tmpl")
	str := s.tRepo.GetMarged(string(b), api)
	path := s.docsDir + name + ".md"
	if s.fRepo.Exist(path) {
		s.fRepo.Remove(path)
	}
	s.fRepo.Write(path, str)
}

// NewDocumenter ... サービスを作成する
func NewDocumenter(
	configDir string,
	docsDir string,
	fRepo repository.File,
	tRepo repository.TemplateClient) Documenter {
	return &documenter{
		configDir: configDir,
		docsDir:   docsDir,
		fRepo:     fRepo,
		tRepo:     tRepo,
	}
}
