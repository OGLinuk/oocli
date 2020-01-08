package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var udugCmd = &cobra.Command{
	Use:   "udug",
	Short: "Update the Onion Omega2",
	Long: `Update the Onion Omega2.
	
	Example Usage

		oocli udug`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Updating ...")
		SSH("opkg update")
	},
}

func init() {
	RootCmd.AddCommand(udugCmd)
}
