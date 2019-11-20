package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rundoc",
	Short: "HTTP/HTTPS準拠のAPI実行 & ドキュメント作成ツール",
	Long:  `HTTP/HTTPS準拠のAPI実行 & ドキュメント作成ツール`,
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
