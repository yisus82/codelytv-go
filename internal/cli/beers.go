package cli

import (
	"bufio"
	"fmt"
	"strconv"

	beerscli "github.com/yisus82/codelytv-go/internal"

	"github.com/spf13/cobra"
)

// CobraBeersFn function definion of run cobra command
type CobraBeersFn func(cmd *cobra.Command, args []string)

const idBeersFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd(repository beerscli.BeerRepo) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(repository),
	}

	beersCmd.Flags().StringP(idBeersFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn(repository beerscli.BeerRepo) CobraBeersFn {
	return func(cmd *cobra.Command, args []string) {
		beers, _ := repository.GetBeers()

		id, _ := cmd.Flags().GetString(idBeersFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			for _, beer := range beers {
				if beer.ProductID == i {
					fmt.Println(beer)
					return
				}
			}
		} else {
			fmt.Println(beers)
		}
	}
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
