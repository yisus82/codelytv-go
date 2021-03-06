package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraStoresFn function definion of run cobra command
type CobraStoresFn func(cmd *cobra.Command, args []string)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

const idStoresFlag = "id"

// InitStoresCmd initialize stores command
func InitStoresCmd() *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   runStoresFn(),
	}

	storesCmd.Flags().StringP(idStoresFlag, "i", "", "id of the store")

	return storesCmd
}

func runStoresFn() CobraStoresFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idStoresFlag)

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
