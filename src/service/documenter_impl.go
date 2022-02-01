package service

import (
	"github.com/aikizoku/rundoc/src/model"
	"github.com/aikizoku/rundoc/src/repository"
)

type documenter struct {
	configDir       string
	docsDir         string
	rFile           repository.File
	rTemplateClient repository.TemplateClient
}

func (s *documenter) Distribute(name string, api *model.API) error {
	b, err := getBinFileData("doc.tmpl")
	if err != nil {
		return err
	}
	str, err := s.rTemplateClient.GetMarged(string(b), api)
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

// NewDocumenter ... サービスを作成する
func NewDocumenter(
	configDir string,
	docsDir string,
	rFile repository.File,
	rTemplateClient repository.TemplateClient) Documenter {
	return &documenter{
		configDir:       configDir,
		docsDir:         docsDir,
		rFile:           rFile,
		rTemplateClient: rTemplateClient,
	}
}
