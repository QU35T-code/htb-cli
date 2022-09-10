package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"os"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a machine",
	Long: `Starts a Hackthebox machine specified in argument`,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetConfigValue("machineid")
		// Active
		url := "https://www.hackthebox.com/api/v4/machine/play/" + machine_id
		var jsonData = []byte("{}")
		resp := utils.HtbPost(url, jsonData)
		message := utils.ParseJsonMessage(resp, "message")
		if (message != "Incorrect lab type.") {
			fmt.Println(message)
			os.Exit(1)
		}
		// Retired
		url = "https://www.hackthebox.com/api/v4/vm/spawn"
		var jsonData2 = []byte(`{"machine_id": ` + machine_id + `}`)
		resp = utils.HtbPost(url, jsonData2)
		message = utils.ParseJsonMessage(resp, "message")
		fmt.Println(message)

		// To Do :
		// Print Machine Name + IP
		
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
