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

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move arg[0] to arg[1] with arg[2] as the name on the Onion Omega2 fileserver",
	Long: `Move arg[0] to arg[1] with arg[2] as the name on the Onion Omega2 fileserver.
	
	This command requires 3 arguments:
		- target | arg[0]
		- destination | arg[1]
		- targets new name | arg[2]
		
	Example Usage

		oocli move hello_world.txt hello+world world_hello.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
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

		mvURL := fmt.Sprintf("http://%s:9001/mv/%s/%s+%s",
			os.Getenv("OOHOST"), args[0], args[1], args[2])

		log.Printf("move URL: %s", mvURL)

		resp, err := client.Get(mvURL)
		if err != nil {
			log.Printf("move.go::move::client.Get(%s)::ERROR: %s", mvURL, err.Error())
		}

		if resp == nil {
			log.Printf("move.go::move::resp::NIL: %s", err.Error())
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			out, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("move.go::move::ioutil.ReadAll(%v)::ERROR: %s", resp.Body, err.Error())
			}
			fmt.Println(string(out))
		} else {
			log.Printf("move.go::move::unknown::ERROR: %s", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(moveCmd)
}
