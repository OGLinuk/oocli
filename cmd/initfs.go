package cmd

import (
	"github.com/spf13/cobra"
)

var initfsCmd = &cobra.Command{
	Use:   "initfs",
	Short: "Initialize the fileserver on the Onion Omega2",
	Long: `Start the Golang fileserver on port 9001. If arg[0]
	is not given, default to ./localfs.
	
	This command requires 0 arguments, BUT has 1 optional:
		- cmd | arg[0]
		
	Example Usage

		oocli initfs ./localfs&`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			args = append(args, "./localfs")
		}

		SSH(args[0])
	},
}

func init() {
	RootCmd.AddCommand(initfsCmd)
}
