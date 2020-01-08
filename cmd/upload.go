package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload arg[0] to arg[1] on the Onion Omega2 fileserver",
	Long: `Upload arg[0] file(s) to the arg[1] given. This will store the file(s)
	in the given place on the Onion Omega2 fileserver. If a destination does not exist,
	one will be created.
	
	The command requires 2 arguments:
		- target | arg[0]
		- destination | arg[1]
		
	Example Usage
	
		oocli upload hello_world.txt hello/world`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		} else if len(args) < 2 {
			args = append(args, ".")
		}
		execCmd := exec.Command("rsync", "--progress", "-azzvh", args[0],
			fmt.Sprintf("%s@%s:/tmp/mounts/SD-P1/%s", os.Getenv("OOUSER"), os.Getenv("OOHOST"), args[1]))

		execCmd.Stderr = os.Stderr
		execCmd.Stdout = os.Stdout

		if err := execCmd.Run(); err != nil {
			log.Printf("upload.go::upload::execCmd.Run()::ERROR: %s", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(uploadCmd)
}
