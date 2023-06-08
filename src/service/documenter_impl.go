package service

import (
	"github.com/aikizoku/rundoc/src/model"
	"github.com/aikizoku/rundoc/src/repository"
)

type documenter struct {
	rFile           repository.File
	rTemplateClient repository.TemplateClient
	configDir       string
	docsDir         string
}

func NewDocumenter(
	rFile repository.File,
	rTemplateClient repository.TemplateClient,
	configDir string,
	docsDir string,
) Documenter {
	return &documenter{
		rFile,
		rTemplateClient,
		configDir,
		docsDir,
	}
}

func (s *documenter) Distribute(
	name string,
	api *model.API,
) error {
	b, err := getBinFileData("doc.tmpl")
	if err != nil {
		return err
	}
	str, err := s.rTemplateClient.GetMerged(string(b), api)
	if err != nil {
		return err
	}
	path := s.docsDir + name + ".md"
	if s.rFile.Exist(path) {
		err = s.rFile.Remove(path)
		if err != nil {
			return err
		}
	}
	err = s.rFile.Write(path, str)
	if err != nil {
		return err
	}
	return nil
}
