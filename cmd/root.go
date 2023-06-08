package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/aikizoku/rundoc/src/config"
	"github.com/aikizoku/rundoc/src/repository"
	"github.com/aikizoku/rundoc/src/service"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	env string
	doc bool
)

var rootCmd = &cobra.Command{
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		if len(args) > 0 {
			name = args[0]
		} else {
			name = ""
		}

		d := &runDependency{}
		d.Inject()

		if name == "" {
			// APIの選択
			names, err := d.Runner.GetRunList()
			if err != nil {
				os.Exit(500)
			}
			index, err := fuzzyfinder.Find(
				names,
				func(i int) string {
					return names[i]
				},
				fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
					if i < 0 {
						return ""
					}
					preview, err := d.Runner.GetRunPreview(names[i])
					if err != nil {
						return err.Error()
					}
					return preview
				}))
			if err != nil {
				os.Exit(404)
			}
			name = names[index]

			// 環境の選択
			envs := []string{"local", "staging", "production"}
			index, err = fuzzyfinder.Find(
				envs,
				func(i int) string {
					return envs[i]
				},
				fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
					if i < 0 {
						return ""
					}
					return "実行環境を選択"
				}))
			if err != nil {
				os.Exit(404)
			}
			env := envs[index]

			// ドキュメント作成の選択
			docs := []bool{false, true}
			index, err = fuzzyfinder.Find(
				docs,
				func(i int) string {
					return fmt.Sprintf("%t", docs[i])
				},
				fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
					if i < 0 {
						return ""
					}
					return "ドキュメントを作成する？"
				}))
			if err != nil {
				os.Exit(404)
			}
			doc := docs[index]

			// 実行
			api, err := d.Runner.Run(name, env, doc)
			if err != nil {
				os.Exit(500)
			}

			if doc {
				err = d.Documenter.Distribute(name, api)
				if err != nil {
					os.Exit(500)
				}
			}

		} else {
			// 自動式
			api, err := d.Runner.Run(name, env, doc)
			if err != nil {
				os.Exit(500)
			}
			if doc {
				err = d.Documenter.Distribute(name, api)
				if err != nil {
					os.Exit(500)
				}
			}
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&env, "env", "e", "local", "APIの実行環境 local, staging, production")
	rootCmd.Flags().BoolVarP(&doc, "doc", "d", false, "実行と同時にドキュメントを作成する")
	env = strings.ToLower(env)
}

type runDependency struct {
	Runner     service.Runner
	Documenter service.Documenter
}

func (d *runDependency) Inject() {
	// Repository
	rFile := repository.NewFile()
	rHTTPClient := repository.NewHTTPClient()
	rTemplateClient := repository.NewTemplateClient()

	// Service
	d.Runner = service.NewRunner(
		rFile,
		rHTTPClient,
		rTemplateClient,
		config.ConfigDir,
		config.RunsDir,
	)
	d.Documenter = service.NewDocumenter(
		rFile,
		rTemplateClient,
		config.ConfigDir,
		config.DocsDir,
	)
}
