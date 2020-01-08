package cmd

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete arg[0] from the Onion Omega2 fileserver",
	Long: `Delete arg[0] from the Onion Omega2 fileserver.
	
	This command requires 1 argument:
		- target | arg[0]
			
	Example Usage

		oocli delete hello_world.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}

		client := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}

		delURL := fmt.Sprintf("http://%s:9001/rm/%s", os.Getenv("OOHOST"), args[0])

		log.Printf("delete URL: %s", delURL)

		resp, err := client.Get(delURL)
		if err != nil {
			log.Printf("delete.go::delete::client.Get(%s)::ERROR: %s", delURL, err.Error())
		}

		if resp == nil {
			log.Printf("delete.go::delete::resp::NIL: %s", err.Error())
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			out, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("delete.go::delete::ioutil.ReadAll(%v)::ERROR: %s", resp.Body, err.Error())
			}
			fmt.Println(string(out))
		} else {
			log.Printf("delete.go::delete::unknown::ERROR: %s", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
