package cmd

import (
	"fmt"
	"io"
	"encoding/json"
	"log"
	"github.com/spf13/cobra"
	"htb/utils"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get Target IP Address",
	Long: `Retrieves the IP address of the machine deployed on Hackthebox and displays it`,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := "492"
		url := "https://www.hackthebox.com/api/v4/machine/profile/" + machine_id
		resp := utils.HtbGet(url)
		json_body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(json_body), &result)
		info := result["info"]
		infomap := info.(map[string]interface{})
    	fmt.Println(infomap["name"])
    	fmt.Println(infomap["ip"])
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
