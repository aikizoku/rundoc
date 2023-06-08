package cmd

import (
	"os"

	"github.com/aikizoku/rundoc/src/config"
	"github.com/aikizoku/rundoc/src/repository"
	"github.com/aikizoku/rundoc/src/service"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "APIリストを表示",
	Long:  `APIリストを表示`,
	Run: func(cmd *cobra.Command, args []string) {
		d := &listDependency{}
		d.Inject()

		err := d.Runner.ShowRunList()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

type listDependency struct {
	Runner service.Runner
}

func (d *listDependency) Inject() {
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
}
