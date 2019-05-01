package main

import (
	"flag"

	"github.com/aikizoku/rundoc/src/config"
	"github.com/aikizoku/rundoc/src/log"
	"github.com/aikizoku/rundoc/src/repository"
	"github.com/aikizoku/rundoc/src/service"
)

func main() {
	isInit := flag.Bool("i", false, "initialize")
	isList := flag.Bool("l", false, "show runs name list")
	name := flag.String("r", "", "runs json file name")
	env := flag.String("e", "local", "run env")
	isDocs := flag.Bool("d", false, "run and generate docs")
	flag.Parse()

	// ログを初期化
	log.Setup()

	// DI
	d := &Dependency{}
	d.Inject()

	// 初期化コマンド
	if *isInit {
		err := d.Initializer.Init()
		if err != nil {
			return
		}
		return
	}

	// 実行リスト表示コマンド
	if *isList {
		err := d.Runner.ShowList()
		if err != nil {
			return
		}
		return
	}

	// 実行コマンド
	api, err := d.Runner.Run(*name, *env)
	if err != nil {
		return
	}
	if *isDocs {
		d.Documenter.Distribute(*name, api)
	}
}

// Dependency ... 依存性
type Dependency struct {
	Initializer service.Initializer
	Runner      service.Runner
	Documenter  service.Documenter
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Repository
	fRepo := repository.NewFile()
	hRepo := repository.NewHTTPClient()
	tRepo := repository.NewTemplateClient()

	// Service
	d.Initializer = service.NewInitializer(config.ConfigDir, config.RunsDir, config.DocsDir, fRepo)
	d.Runner = service.NewRunner(config.ConfigDir, config.RunsDir, fRepo, hRepo, tRepo)
	d.Documenter = service.NewDocumenter(config.ConfigDir, config.DocsDir, fRepo, tRepo)
}
