package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"

	"github.com/aikizoku/rundoc/src/config"
	"github.com/aikizoku/rundoc/src/repository"
	"github.com/aikizoku/rundoc/src/service"
)

var (
	env string
	doc bool
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "APIを実行",
	Long:  `APIを実行`,
	Args:  cobra.MaximumNArgs(1),
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
				os.Exit(1)
			}
			index, err := fuzzyfinder.Find(
				names,
				func(i int) string {
					return names[i]
				},
				fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
					preview, err := d.Runner.GetRunPreview(names[i])
					if err != nil {
						return err.Error()
					}
					return preview
				}))
			name = names[index]

			// 環境の選択
			envs := []string{"local", "staging", "production"}
			index, err = fuzzyfinder.Find(
				envs,
				func(i int) string {
					return envs[i]
				},
				fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
					return "実行環境を選択"
				}))
			env := envs[index]

			// ドキュメント作成の選択
			docs := []bool{false, true}
			index, err = fuzzyfinder.Find(
				docs,
				func(i int) string {
					return fmt.Sprintf("%t", docs[i])
				},
				fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
					return "ドキュメントを作成する？"
				}))
			doc := docs[index]

			// 実行
			api, err := d.Runner.Run(name, env, doc)
			if err != nil {
				os.Exit(1)
			}

			if doc {
				d.Documenter.Distribute(name, api)
			}

		} else {
			// 自動式
			api, err := d.Runner.Run(name, env, doc)
			if err != nil {
				os.Exit(1)
			}
			if doc {
				d.Documenter.Distribute(name, api)
			}
		}
	},
}

func init() {
	runCmd.Flags().StringVarP(&env, "env", "e", "local", "APIの実行環境 local, staging, production")
	runCmd.Flags().BoolVarP(&doc, "doc", "d", false, "実行と同時にドキュメントを作成する")
	env = strings.ToLower(env)
	rootCmd.AddCommand(runCmd)
}

type runDependency struct {
	Runner     service.Runner
	Documenter service.Documenter
}

func (d *runDependency) Inject() {
	// Repository
	fRepo := repository.NewFile()
	hRepo := repository.NewHTTPClient()
	tRepo := repository.NewTemplateClient()

	// Service
	d.Runner = service.NewRunner(config.ConfigDir, config.RunsDir, fRepo, hRepo, tRepo)
	d.Documenter = service.NewDocumenter(config.ConfigDir, config.DocsDir, fRepo, tRepo)
}
