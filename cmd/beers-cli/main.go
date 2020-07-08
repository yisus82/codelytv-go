package main

import (
	"github.com/spf13/cobra"
	"github.com/yisus82/codelytv-go/internal/cli"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}
