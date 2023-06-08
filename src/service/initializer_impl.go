package service

import (
	"encoding/json"

	"github.com/aikizoku/rundoc/src/log"
	"github.com/aikizoku/rundoc/src/model"
	"github.com/aikizoku/rundoc/src/repository"
)

type initializer struct {
	rFile     repository.File
	rootDir   string
	configDir string
	runsDir   string
	docsDir   string
}

func NewInitializer(
	rFile repository.File,
	rootDir string,
	configDir string,
	runsDir string,
	docsDir string,
) Initializer {
	return &initializer{
		rFile,
		rootDir,
		configDir,
		runsDir,
		docsDir,
	}
}

func (s *initializer) Init() error {
	if err := s.rFile.WriteDir(s.rootDir); err != nil {
		panic(err)
	}
	if err := s.rFile.WriteDir(s.configDir); err != nil {
		panic(err)
	}
	if err := s.rFile.WriteDir(s.runsDir); err != nil {
		panic(err)
	}
	if err := s.rFile.WriteDir(s.docsDir); err != nil {
		panic(err)
	}

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
		jStr, err := convertPrettyJSON(j, false)
		if err != nil {
			return err
		}
		err = s.rFile.Write(s.configDir+"common.json", jStr)
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
		jStr, err := convertPrettyJSON(j, false)
		if err != nil {
			return err
		}
		err = s.rFile.Write(s.configDir+"auth.json", jStr)
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
			Params: map[string]any{
				"hoge": "aaaaa",
				"fuga": "xxxxx",
			},
		}
		j, err := json.Marshal(sample)
		if err != nil {
			log.Errorf(err, "sample.jsonのparseに失敗: %v", sample)
			return err
		}
		jStr, err := convertPrettyJSON(j, false)
		if err != nil {
			return err
		}
		err = s.rFile.Write(s.runsDir+"sample.json", jStr)
		if err != nil {
			return err
		}
	}
	return nil
}
