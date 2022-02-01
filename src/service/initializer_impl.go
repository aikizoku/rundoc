package service

import (
	"encoding/json"

	"github.com/aikizoku/rundoc/src/log"
	"github.com/aikizoku/rundoc/src/model"
	"github.com/aikizoku/rundoc/src/repository"
)

type initializer struct {
	rootDir   string
	configDir string
	runsDir   string
	docsDir   string
	rFile     repository.File
}

func (s *initializer) Init() error {
	s.rFile.WriteDir(s.rootDir)
	s.rFile.WriteDir(s.configDir)
	s.rFile.WriteDir(s.runsDir)
	s.rFile.WriteDir(s.docsDir)

	// common.json
	if !s.rFile.Exist(s.configDir + "common.json") {
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
			log.Errorf(err, "common.jsonのparseに失敗: %v", common)
			return err
		}
		jstr, err := convertPrettyJSON(j)
		if err != nil {
			return err
		}
		err = s.rFile.Write(s.configDir+"common.json", jstr)
		if err != nil {
			return err
		}
		return nil
	}

	// auth.json
	if !s.rFile.Exist(s.configDir + "auth.json") {
		auth := &model.FileAuth{
			Local:      "sample_local_token",
			Staging:    "sample_staging_token",
			Production: "sample_production_token",
		}
		j, err := json.Marshal(auth)
		if err != nil {
			log.Errorf(err, "auth.jsonのparseに失敗: %v", auth)
			return err
		}
		jstr, err := convertPrettyJSON(j)
		if err != nil {
			return err
		}
		err = s.rFile.Write(s.configDir+"auth.json", jstr)
		if err != nil {
			return err
		}
	}

	// sample.json
	if !s.rFile.Exist(s.runsDir + "sample.json") {
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
			log.Errorf(err, "sample.jsonのparseに失敗: %v", sample)
			return err
		}
		jstr, err := convertPrettyJSON(j)
		if err != nil {
			return err
		}
		err = s.rFile.Write(s.runsDir+"sample.json", jstr)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewInitializer ...
func NewInitializer(
	rootDir string,
	configDir string,
	runsDir string,
	docsDir string,
	rFile repository.File) Initializer {
	return &initializer{
		rootDir:   rootDir,
		configDir: configDir,
		runsDir:   runsDir,
		docsDir:   docsDir,
		rFile:     rFile,
	}
}
