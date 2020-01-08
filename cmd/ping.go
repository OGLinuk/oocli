package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the Onion Omega2",
	Long: `Ping the Onion Omega2.
			
	Example Usage

		oocli ping`,
	Run: func(cmd *cobra.Command, args []string) {
		pingOOCMD := exec.Command("ping", os.Getenv("OOHOST"))

		pingOOCMD.Stderr = os.Stderr
		pingOOCMD.Stdout = os.Stdout

		if err := pingOOCMD.Run(); err != nil {
			log.Printf("ping.go::pingCmd::pingOOCMD.Run()::ERROR: %s", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(pingCmd)
}
