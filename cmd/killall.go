package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var killallCmd = &cobra.Command{
	Use:   "killall",
	Short: "Execute killall arg[0]",
	Long: `Allows for the killall command to be executed with the
	given args[0].
	
	This command requires 1 argument:
		- target | arg[0]
		
	Example Usage

		oocli killall localfs`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}

		SSH(fmt.Sprintf("killall %s", args[0]))
	},
}

func init() {
	RootCmd.AddCommand(killallCmd)
}
