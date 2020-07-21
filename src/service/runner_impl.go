package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/aikizoku/rundoc/src/log"
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

func (s *runner) ShowRunList() error {
	fileNames, err := s.fRepo.GetNameList(s.runsDir)
	if err != nil {
		return err
	}
	for _, fileName := range fileNames {
		name := s.getFileNameWithoutExt(fileName)
		fmt.Println(name)
	}
	return nil
}

func (s *runner) GetRunList() ([]string, error) {
	fileNames, err := s.fRepo.GetNameList(s.runsDir)
	if err != nil {
		return nil, err
	}
	dsts := []string{}
	for _, fileName := range fileNames {
		dst := s.getFileNameWithoutExt(fileName)
		dsts = append(dsts, dst)
	}
	return dsts, nil
}

func (s *runner) getFileNameWithoutExt(fileName string) string {
	return filepath.Base(fileName[:len(fileName)-len(filepath.Ext(fileName))])
}

func (s *runner) GetRunPreview(name string) (string, error) {
	// 実行設定
	run, err := s.getRunFile(name)
	if err != nil {
		log.Errorf(err, "ファイル読み込みに失敗: %s%s%s", s.runsDir, name, ".json")
		return "", err
	}

	// Header文字列を作成
	hStrs := []string{}
	if _, ok := run.Headers["Authorization"]; !ok {
		// 認証設定
		auth, err := s.getAuthFile()
		if err != nil {
			log.Errorf(err, "ファイル読み込みに失敗: %s%s", s.configDir, "auth.json")
			return "", err
		}
		hStrs = append(hStrs, fmt.Sprintf("Authorization(Local): %s", auth.Local))
		hStrs = append(hStrs, fmt.Sprintf("Authorization(Staging): %s", auth.Staging))
		hStrs = append(hStrs, fmt.Sprintf("Authorization(Production): %s", auth.Production))
	}
	hStrs = append(hStrs, "Content-Type: application/json")
	for key, value := range run.Headers {
		hStrs = append(hStrs, fmt.Sprintf("%s: %s", key, value))
	}

	// Param文字列を作成
	params, err := json.Marshal(run.Params)
	if err != nil {
		log.Errorf(err, "jsonのparseに失敗: %v", run.Params)
		return "", err
	}
	pStr, err := convertPrettyJSON(params)
	if err != nil {
		return "", err
	}

	// APIを作成
	api := &model.API{}
	api.Name = name
	api.Description = run.Description
	api.Request = &model.APIRequest{
		Method:  strings.ToUpper(run.Method),
		Path:    run.Path,
		Headers: strings.Join(hStrs, "\n"),
		Params:  strings.Trim(pStr, "\n"),
	}

	// プレビュー
	b, err := getBinFileData("print_preview.tmpl")
	if err != nil {
		return "", err
	}
	out, err := s.tRepo.GetMarged(string(b), api)
	if err != nil {
		return "", err
	}
	return out, nil
}

func (s *runner) getCommonFile() (*model.FileCommon, error) {
	file, err := ioutil.ReadFile(s.configDir + "common.json")
	if err != nil {
		log.Errorf(err, "ファイル読み込みに失敗: %s%s", s.configDir, "common.json")
		return nil, err
	}
	var dst model.FileCommon
	err = json.Unmarshal(file, &dst)
	if err != nil {
		log.Errorf(err, "jsonのparseに失敗: %s", string(file))
		return nil, err
	}
	return &dst, nil
}

func (s *runner) getAuthFile() (*model.FileAuth, error) {
	file, err := ioutil.ReadFile(s.configDir + "auth.json")
	if err != nil {
		log.Errorf(err, "ファイル読み込みに失敗: %s%s", s.configDir, "common.json")
		return nil, err
	}
	var dst model.FileAuth
	err = json.Unmarshal(file, &dst)
	if err != nil {
		log.Errorf(err, "jsonのparseに失敗: %s", string(file))
		return nil, err
	}
	return &dst, nil
}

func (s *runner) getRunFile(name string) (*model.FileRun, error) {
	file, err := ioutil.ReadFile(s.runsDir + name + ".json")
	if err != nil {
		log.Errorf(err, "ファイル読み込みに失敗: %s%s%s", s.runsDir, name, ".json")
		return nil, err
	}
	var dst model.FileRun
	err = json.Unmarshal(file, &dst)
	if err != nil {
		log.Errorf(err, "jsonのparseに失敗: %s", string(file))
		return nil, err
	}
	dst.Method = strings.ToLower(dst.Method)
	return &dst, nil
}

func (s *runner) Run(name string, env string, doc bool) (*model.API, error) {
	// 共通設定
	common, err := s.getCommonFile()
	if err != nil {
		log.Errorf(err, "ファイル読み込みに失敗: %s%s", s.configDir, "common.json")
		return nil, err
	}

	// 認証設定
	auth, err := s.getAuthFile()
	if err != nil {
		log.Errorf(err, "ファイル読み込みに失敗: %s%s", s.configDir, "auth.json")
		return nil, err
	}

	// 実行設定
	run, err := s.getRunFile(name)
	if err != nil {
		log.Errorf(err, "ファイル読み込みに失敗: %s%s%s", s.runsDir, name, ".json")
		return nil, err
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
		err = fmt.Errorf("invalid env: %s", env)
	}
	if err != nil {
		log.Errorf(err, "不正なenv: %s", env)
		return nil, err
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
		log.Errorf(err, "jsonのparseに失敗: %v", run.Params)
		return nil, err
	}

	// 実行
	var runTime int64
	var statusCode int
	var body []byte
	switch run.Method {
	case "get":
		runTime, statusCode, body, err = s.hRepo.Get(url, run.Params, headers)
		if err != nil {
			return nil, err
		}
	case "post":
		runTime, statusCode, body, err = s.hRepo.Post(url, params, headers)
		if err != nil {
			return nil, err
		}
	case "put":
		runTime, statusCode, body, err = s.hRepo.Put(url, params, headers)
		if err != nil {
			return nil, err
		}
	case "delete":
		runTime, statusCode, body, err = s.hRepo.Delete(url, run.Params, headers)
		if err != nil {
			return nil, err
		}
	default:
		err = fmt.Errorf("invalid method: %s", run.Method)
		log.Errorf(err, "不正なmethod: %s", run.Method)
		return nil, err
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
	reqStr, err := convertPrettyJSON(params)
	if err != nil {
		return nil, err
	}
	api.Request = &model.APIRequest{
		Method:  strings.ToUpper(run.Method),
		Path:    run.Path,
		Headers: strings.Join(ahStrs, "\n"),
		Params:  strings.Trim(reqStr, "\n"),
	}
	resStr, err := convertPrettyJSON(body)
	if err != nil {
		return nil, err
	}
	api.Response = &model.APIResponse{
		Time:       runTime,
		StatusCode: statusCode,
		Body:       strings.Trim(resStr, "\n"),
	}
	api.Command = s.generateCommand(name, env, doc)

	// 結果(リクエスト)を表示
	b, err := getBinFileData("print_req.tmpl")
	if err != nil {
		return nil, err
	}
	out, err := s.tRepo.GetMarged(string(b), api)
	if err != nil {
		return nil, err
	}
	fmt.Println("\n\x1b[32m" + out + "\x1b[0m")

	// 結果(レスポンス)を表示
	b, err = getBinFileData("print_res.tmpl")
	if err != nil {
		return nil, err
	}
	out, err = s.tRepo.GetMarged(string(b), api)
	if err != nil {
		return nil, err
	}
	fmt.Println("\n\x1b[36m" + out + "\x1b[0m")

	// コマンド
	b, err = getBinFileData("print_cmd.tmpl")
	if err != nil {
		return nil, err
	}
	out, err = s.tRepo.GetMarged(string(b), api)
	if err != nil {
		return nil, err
	}
	fmt.Println("\n\x1b[35m" + out + "\x1b[0m")
	fmt.Println("")

	// 認証情報を隠したheaderに差し替える
	api.Request.Headers = strings.Join(hStrs, "\n")

	return api, nil
}

func (s *runner) generateCommand(name string, env string, doc bool) string {
	cmds := []string{"rundoc", "run", name}
	switch env {
	case "local":
		break
	case "staging", "production":
		cmds = append(cmds, "-e")
		cmds = append(cmds, env)
	}
	if doc {
		cmds = append(cmds, "-d")
	}
	return strings.Join(cmds, " ")
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
