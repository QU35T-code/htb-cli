package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"encoding/json"
	"io"
	"log"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a machine",
	Long: `Starts a Hackthebox machine specified in argument`,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := "492" // Temp
		url := "https://www.hackthebox.com/api/v4/machine/play/" + machine_id
		var jsonData = []byte("{}")
		resp := utils.HtbPost(url, jsonData)
		json_body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(json_body), &result)
    	fmt.Println(result["message"])

		// To Do :
		// Print Machine Name + IP
		
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
