package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current machine",
	Long: "Stop the current machine",
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetConfigValue("machineid")
		url := "https://www.hackthebox.com/api/v4/machine/stop"
		var jsonData = []byte(`{"machine_id": ` + machine_id + `}`)
		resp := utils.HtbPost(url, jsonData)
		message := utils.ParseJsonMessage(resp, "message")
		fmt.Println("Active Machines")
		fmt.Println(message)

		// Retired
		url = "https://www.hackthebox.com/api/v4/vm/terminate"
		var jsonData2 = []byte(`{"machine_id": ` + machine_id + `}`)
		resp = utils.HtbPost(url, jsonData2)
		message = utils.ParseJsonMessage(resp, "message")
		fmt.Println("Retired Machines")
		fmt.Println(message)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
