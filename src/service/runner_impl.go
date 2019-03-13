package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/aikizoku/rundoc/src/model"
	"github.com/aikizoku/rundoc/src/repository"
)

type runner struct {
	configDir string
	runsDir   string
	fRepo     repository.File
	hRepo     repository.HTTPClient
	tRepo     repository.TemplateClient
}

func (s *runner) ShowList() {
	fileNames := s.fRepo.GetNameList(s.runsDir)

	fmt.Println("----- runable names -----")
	for _, fileName := range fileNames {
		name := s.getFileNameWithoutExt(fileName)
		fmt.Println(name)
	}
	fmt.Println("-------------------------")
	return
}

func (s *runner) getFileNameWithoutExt(fileName string) string {
	return filepath.Base(fileName[:len(fileName)-len(filepath.Ext(fileName))])
}

func (s *runner) Run(name string, env string) *model.API {
	// 共通設定
	commonFile, err := ioutil.ReadFile(s.configDir + "common.json")
	if err != nil {
		panic(err)
	}
	var common model.FileCommon
	err = json.Unmarshal(commonFile, &common)
	if err != nil {
		panic(err)
	}

	// 認証設定
	authFile, err := ioutil.ReadFile(s.configDir + "auth.json")
	if err != nil {
		panic(err)
	}
	var auth model.FileAuth
	err = json.Unmarshal(authFile, &auth)
	if err != nil {
		panic(err)
	}

	// 実行設定
	runFile, err := ioutil.ReadFile(s.runsDir + name + ".json")
	if err != nil {
		panic(err)
	}
	var run model.FileRun
	err = json.Unmarshal(runFile, &run)
	if err != nil {
		panic(err)
	}

	// 環境選択
	var url, authorization string
	switch env {
	case "local":
		url = common.Endpoints.Local + run.Path
		authorization = auth.Local
	case "staging":
		url = common.Endpoints.Staging + run.Path
		authorization = auth.Staging
	case "production":
		url = common.Endpoints.Production + run.Path
		authorization = auth.Production
	default:
		panic(fmt.Errorf("invalid env: %s", env))
	}

	// Header結合
	headers := map[string]string{}
	headers["Authorization"] = authorization
	for key, value := range common.Headers {
		headers[key] = value
	}
	for key, value := range run.Headers {
		headers[key] = value
	}

	// Params
	params, err := json.Marshal(run.Params)
	if err != nil {
		panic(err)
	}

	// 実行
	var runTime int64
	var statusCode int
	var body []byte
	switch run.Method {
	case "get":
		runTime, statusCode, body = s.hRepo.Get(url, run.Params, headers)
	case "post":
		runTime, statusCode, body = s.hRepo.Post(url, params, headers)
	case "put":
		runTime, statusCode, body = s.hRepo.Put(url, params, headers)
	case "delete":
		runTime, statusCode, body = s.hRepo.Delete(url, run.Params, headers)
	default:
		panic(fmt.Errorf("invalid method: %s", run.Method))
	}

	// header文字列を作成
	hStrs := []string{}
	ahStrs := []string{}
	for key, value := range headers {
		var hStr, ahStr string

		// 認証情報を隠したheader文字列を作成
		if key == "Authorization" {
			hStr = fmt.Sprintf("%s: %s", key, "xxxxxxxxxx")
		} else {
			hStr = fmt.Sprintf("%s: %s", key, value)
		}
		hStrs = append(hStrs, hStr)

		// 認証情報があるheader文字列を作成
		ahStr = fmt.Sprintf("%s: %s", key, value)
		ahStrs = append(ahStrs, ahStr)
	}

	// 結果を整理
	api := &model.API{}
	api.Name = name
	api.Description = run.Description
	api.Endpoints = &model.APIEndpoints{
		Local:      common.Endpoints.Local,
		Staging:    common.Endpoints.Staging,
		Production: common.Endpoints.Production,
	}
	api.Request = &model.APIRequest{
		Method:  strings.ToUpper(run.Method),
		Path:    run.Path,
		Headers: strings.Join(ahStrs, "\n"),
		Params:  strings.Trim(convertPrettyJSON(params), "\n"),
	}
	api.Response = &model.APIResponse{
		Time:       runTime,
		StatusCode: statusCode,
		Body:       strings.Trim(convertPrettyJSON(body), "\n"),
	}

	// 結果を表示
	b := getBinFileData("print.tmpl")
	out := s.tRepo.GetMarged(string(b), api)
	fmt.Println(out)

	// 認証情報を隠したheaderに差し替える
	api.Request.Headers = strings.Join(hStrs, "\n")

	return api
}

// NewRunner ... サービスを作成する
func NewRunner(
	configDir string,
	runsDir string,
	fRepo repository.File,
	hRepo repository.HTTPClient,
	tRepo repository.TemplateClient) Runner {
	return &runner{
		configDir: configDir,
		runsDir:   runsDir,
		fRepo:     fRepo,
		hRepo:     hRepo,
		tRepo:     tRepo,
	}
}
