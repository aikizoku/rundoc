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

func (s *documenter) Distribute(name string, api *model.API) error {
	b, err := getBinFileData("doc.tmpl")
	if err != nil {
		return err
	}
	str, err := s.tRepo.GetMarged(string(b), api)
	if err != nil {
		return err
	}
	path := s.docsDir + name + ".md"
	if s.fRepo.Exist(path) {
		err = s.fRepo.Remove(path)
		if err != nil {
			return err
		}
	}
	err = s.fRepo.Write(path, str)
	if err != nil {
		return err
	}
	return nil
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
