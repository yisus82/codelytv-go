package main

import (
	"github.com/spf13/cobra"
	"github.com/yisus82/codelytv-go/internal/cli"
	"github.com/yisus82/codelytv-go/internal/storage/csv"
)

func main() {
	csvRepo := csv.NewRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(csvRepo))
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}
