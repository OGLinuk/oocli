package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "oocli",
	Short: "A CLI to interface with an Onion Omega2 fileserver ...",
	Long: `The Onion Omega2 runs OpenWRT, which allows the device to act as an access
	point and bridge WiFi to connected devices. Since it runs BusyBox (Linux), it can 
	run a small Golang cross-compiled fileserver. Any device connected to the Onion 
	Omega2 will have access to the fileserver via 192.168.3.1:9001. This CLI is used 
	to easily interact with that fileserver. The media files are also available to be 
	viewed, and all files to be downloaded through firefox.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	print("\033[H\033[2J") //clear terminal

	fmt.Println(`
4F 6E 69 6F 6E  4F 6D 65 67 61 32  43 4C 49

╔═╗┌┐┌┬┌─┐┌┐┌  ╔═╗┌┬┐┌─┐┌─┐┌─┐ ┬┬  ╔═╗╦  ╦
║ ║│││││ ││││  ║ ║│││├┤ │ ┬├─┤ ││  ║  ║  ║
╚═╝┘└┘┴└─┘┘└┘  ╚═╝┴ ┴└─┘└─┘┴ ┴ ┴┴  ╚═╝╩═╝╩
`)
}
