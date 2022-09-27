package cmd

import (
	"fmt"
	"os"

	"github.com/QU35T-code/htb-cli/utils"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a machine",
	Long:  "Reset a machine",
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := os.Getenv("HTB_MACHINE_ID")
		machine_type := utils.GetMachineType(machine_id)
		if machine_type == "active" {
			url := "https://www.hackthebox.com/api/v4/vm/reset"
			var jsonData = []byte(`{"machine_id": ` + machine_id + `}`)
			resp := utils.HtbPost(url, jsonData)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Println(message)
		} else {
			url := "https://www.hackthebox.com/api/v4/release_arena/reset"
			var jsonData = []byte(`{"machine_id": ` + machine_id + `}`)
			resp := utils.HtbPost(url, jsonData)
			message := utils.ParseJsonMessage(resp, "message")
			fmt.Println(message)
		}

	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
