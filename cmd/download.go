package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download arg[0] to arg[1] from the Onion Omega2 fileserver",
	Long: `Download arg[0] from the Onion Omega2 fileserver to the given arg[1] path.
		
	This command requires 2 arguments:
		- target | arg[0]
		- destination | arg[1]
		
	Example Usage
	
		oocli download hello_world.txt .`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		} else if len(args) < 2 {
			args = append(args, ".")
		}

		execCmd := exec.Command("rsync", "--progress", "-azzvh",
			fmt.Sprintf("%s@%s:/tmp/mounts/SD-P1/%s", os.Getenv("OOUSER"), os.Getenv("OOHOST"), args[0]), args[1])

		execCmd.Stderr = os.Stderr
		execCmd.Stdout = os.Stdout

		if err := execCmd.Run(); err != nil {
			log.Printf("download.go::download::execCmd.Run()::ERROR: %s", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(downloadCmd)
}
