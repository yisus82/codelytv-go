package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// CobraBeersFn function definion of run cobra command
type CobraBeersFn func(cmd *cobra.Command, args []string)

const idBeersFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(),
	}

	beersCmd.Flags().StringP(idBeersFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn() CobraBeersFn {
	return func(cmd *cobra.Command, args []string) {
		f, _ := os.Open("../../data/beers.csv")
		reader := bufio.NewReader(f)

		beers := make(map[int]string)

		for line := readLine(reader); line != nil; line = readLine(reader) {
			values := strings.Split(string(line), ",")

			productID, _ := strconv.Atoi(values[0])
			beers[productID] = values[1]
		}

		id, _ := cmd.Flags().GetString(idBeersFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			fmt.Println(beers[i])
		} else {
			fmt.Println(beers)
		}
	}
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
