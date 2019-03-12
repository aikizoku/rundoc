package service

import (
	"github.com/aikizoku/rundoc/src/model"
	"github.com/aikizoku/rundoc/src/repository"
)

type doc struct {
	configDir string
	docsDir   string
	fRepo     repository.File
	tRepo     repository.TemplateClient
}

func (s *doc) Distribute(name string, api *model.API) {
	str := s.tRepo.GetMarged(s.configDir+"doc.tmpl", api)
	fileName := name + ".md"
	if s.fRepo.Exist(s.docsDir, fileName) {
		s.fRepo.Remove(s.docsDir, fileName)
	}
	s.fRepo.Write(s.docsDir, fileName, str)
}

// NewDoc ... サービスを作成する
func NewDoc(configDir string, docsDir string, fRepo repository.File, tRepo repository.TemplateClient) Doc {
	return &doc{
		configDir: configDir,
		docsDir:   docsDir,
		fRepo:     fRepo,
		tRepo:     tRepo,
	}
}
