package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.GetConfigValue("machineid")
		url := "https://www.hackthebox.com/api/v4/machine/profile/" + machine_id
		resp := utils.HtbGet(url)
		info := utils.ParseJsonMessage(resp, "info").(map[string]interface{})
		fmt.Println("--- INFO ---")
		fmt.Print("Name : ")
		fmt.Println(info["name"])
		fmt.Print("OS : ")
		fmt.Println(info["os"])
		fmt.Print("Points : ")
		fmt.Println(info["points"])
		fmt.Print("Difficulty : ")
		fmt.Println(info["difficultyText"])
		fmt.Print("Is Completed ? ")
		fmt.Println(info["isCompleted"])
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
