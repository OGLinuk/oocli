package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the Onion Omega2",
	Long: `Restart the Onion Omega2.
			
	Example Usage

		oocli restart`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("restarting the Onion Omega2 ...")
		SSH("reboot")
	},
}

func init() {
	RootCmd.AddCommand(restartCmd)
}
