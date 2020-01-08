package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var psgrepCmd = &cobra.Command{
	Use:   "psgrep",
	Short: "Execute 'ps | grep arg[0]' on the Onion Omega2",
	Long: `Execute 'ps | grep arg[0]' on the Onion Omega2.
	
	This command requires 1 argument:
		- filter | arg[0]
			
	Example Usage

		oocli psgrep localfs`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}

		SSH(fmt.Sprintf("ps | grep %s", args[0]))
	},
}

func init() {
	RootCmd.AddCommand(psgrepCmd)
}
