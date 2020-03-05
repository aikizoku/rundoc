package cmd

import (
	"os"

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
			err := d.Runner.ShowList()
			if err != nil {
				os.Exit(1)
			}
		} else {
			api, err := d.Runner.Run(name, env)
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
