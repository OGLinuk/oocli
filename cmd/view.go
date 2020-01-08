package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View arg[0] from the Onion Omega2 fileserver on the default browser",
	Long: `View opens the arg[0] in a firefox browser, which has the
	built-in functionality to view media like pdf and mp4 formatted files.
	
		
	This command requires 0 arguments, BUT has 1 optional:
		- target | arg[0]
			
	Example Usage

		oocli view hello.pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Opening browser ...")
		open.Start(fmt.Sprintf("http://%s:9001", os.Getenv("OOHOST")))
	},
}

func init() {
	RootCmd.AddCommand(viewCmd)
}
