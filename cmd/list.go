package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/aikizoku/rundoc/src/config"
	"github.com/aikizoku/rundoc/src/repository"
	"github.com/aikizoku/rundoc/src/service"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "APIリストを表示",
	Long:  `APIリストを表示`,
	Run: func(cmd *cobra.Command, args []string) {
		d := &listDependency{}
		d.Inject()

		err := d.Runner.ShowList()
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
	fRepo := repository.NewFile()
	hRepo := repository.NewHTTPClient()
	tRepo := repository.NewTemplateClient()

	// Service
	d.Runner = service.NewRunner(config.ConfigDir, config.RunsDir, fRepo, hRepo, tRepo)
}
