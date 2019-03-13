package service

import (
	"encoding/json"

	"github.com/aikizoku/rundoc/src/model"
	"github.com/aikizoku/rundoc/src/repository"
)

type initializer struct {
	configDir string
	runsDir   string
	docsDir   string
	fRepo     repository.File
}

func (s *initializer) Init() {
	s.fRepo.WriteDir(s.configDir)
	s.fRepo.WriteDir(s.runsDir)
	s.fRepo.WriteDir(s.docsDir)

	// common.json
	if !s.fRepo.Exist(s.configDir + "common.json") {
		common := &model.FileCommon{
			Endpoints: &model.FileEndpoints{
				Local:      "http://localhost:8080",
				Staging:    "https://staging.appspot.com",
				Production: "https://appspot.com",
			},
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		j, err := json.Marshal(common)
		if err != nil {
			panic(err)
		}
		s.fRepo.Write(s.configDir+"common.json", convertPrettyJSON(j))
	}

	// auth.json
	if !s.fRepo.Exist(s.configDir + "auth.json") {
		auth := &model.FileAuth{
			Local:      "sample_local_token",
			Staging:    "sample_staging_token",
			Production: "sample_production_token",
		}
		j, err := json.Marshal(auth)
		if err != nil {
			panic(err)
		}
		s.fRepo.Write(s.configDir+"auth.json", convertPrettyJSON(j))
	}

	// sample.json
	if !s.fRepo.Exist(s.runsDir + "sample.json") {
		sample := &model.FileRun{
			Description: "サンプルAPIの詳細",
			Path:        "/v1/sample",
			Method:      "post",
			Headers: map[string]string{
				"X-OS": "iOS",
			},
			Params: map[string]interface{}{
				"hoge": "aaaaa",
				"fuga": "xxxxx",
			},
		}
		j, err := json.Marshal(sample)
		if err != nil {
			panic(err)
		}
		s.fRepo.Write(s.runsDir+"sample.json", convertPrettyJSON(j))
	}
}

// NewInitializer ...
func NewInitializer(
	configDir string,
	runsDir string,
	docsDir string,
	fRepo repository.File) Initializer {
	return &initializer{
		configDir: configDir,
		runsDir:   runsDir,
		docsDir:   docsDir,
		fRepo:     fRepo,
	}
}
