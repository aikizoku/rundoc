package main

import (
	"flag"

	"github.com/aikizoku/rundoc/src/config"
	"github.com/aikizoku/rundoc/src/repository"
	"github.com/aikizoku/rundoc/src/service"
)

func main() {
	name := flag.String("n", "", "runs json file name")
	isList := flag.Bool("l", false, "show runs name list")
	env := flag.String("e", "local", "run env")
	isDocs := flag.Bool("d", false, "run and generate docs")
	flag.Parse()

	d := &Dependency{}
	d.Inject()

	if *isList {
		d.Run.ShowList()
	} else {
		api := d.Run.Run(*name, *env)
		if *isDocs {
			d.Doc.Distribute(*name, api)
		}
	}
}

// Dependency ... 依存性
type Dependency struct {
	Run service.Run
	Doc service.Doc
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Repository
	fRepo := repository.NewFile()
	hRepo := repository.NewHTTPClient()
	tRepo := repository.NewTemplateClient()

	// Service
	d.Run = service.NewRun(config.ConfigDir, config.RunsDir, fRepo, hRepo, tRepo)
	d.Doc = service.NewDoc(config.ConfigDir, config.DocsDir, fRepo, tRepo)
}
