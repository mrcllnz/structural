package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/structural/cmd/st/ga"
)

var hackLong = `dev command`

func HackRun(args []string) (err error) {

	return err
}

var HackCmd = &cobra.Command{

	Use: "hack",

	Hidden: true,

	Short: "dev command",

	Long: hackLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, strings.Join(args, "/"), 0)

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = HackRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
