package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Displays machine information",
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetActiveMachineID()
		machine_id = fmt.Sprintf("%v", machine_id)
		url := "https://www.hackthebox.com/api/v4/machine/profile/" + machine_id.(string)
		resp := utils.HtbGet(url)
		info := utils.ParseJsonMessage(resp, "info").(map[string]interface{})
		fmt.Printf("--- INFO ---\nName : %v\nOS : %v\nPoints : %v\nDifficulty : %v\nIs Completed ? %v", info["name"], info["os"], info["points"], info["difficultyText"], info["isCompleted"])
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
