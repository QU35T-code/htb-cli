package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"encoding/json"
	"io"
	"log"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current machine",
	Long: "Stop the current machine",
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://www.hackthebox.com/api/v4/machine/stop"
		var jsonData = []byte("{}")
		resp := utils.HtbPost(url, jsonData)
		json_body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(json_body), &result)
    	fmt.Println(result["message"])
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
