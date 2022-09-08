package cmd

import (
	"fmt"
	"io"
	"encoding/json"
	"log"
	"net/http"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get Target IP Address",
	Long: `Retrieves the IP address of the machine deployed on Hackthebox and displays it`,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := "492"
		url := "https://www.hackthebox.com/api/v4/machine/profile/" + machine_id
		jwt_token := "REDACTED"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("User-Agent", "HTB-Tool")
		req.Header.Set("Authorization", "Bearer " + jwt_token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
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
