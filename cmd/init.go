package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/aikizoku/rundoc/src/config"
	"github.com/aikizoku/rundoc/src/repository"
	"github.com/aikizoku/rundoc/src/service"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "rundocの実行環境を作成",
	Long: `runsディレクトリの作成
docsディレクトリの作成
configディレクトリの作成`,
	Run: func(cmd *cobra.Command, args []string) {
		d := &initDependency{}
		d.Inject()

		err := d.Initializer.Init()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

type initDependency struct {
	Initializer service.Initializer
}

func (d *initDependency) Inject() {
	// Repository
	rFile := repository.NewFile()

	// Service
	d.Initializer = service.NewInitializer(config.RootDir, config.ConfigDir, config.RunsDir, config.DocsDir, rFile)
}
