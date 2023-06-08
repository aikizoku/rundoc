package cmd

import (
	"fmt"

	"github.com/aikizoku/rundoc/src/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "現在のバージョンを表示する",
	Long:  `現在のバージョンを表示する`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("version %s\n", config.AppVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
