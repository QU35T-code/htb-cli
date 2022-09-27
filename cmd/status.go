package cmd

import (
	"fmt"
	"htb-cli/utils"
	"os"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Displays the active machine",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://www.hackthebox.com/api/v4/machine/active"
		resp := utils.HtbGet(url)
		info := utils.ParseJsonMessage(resp, "info")
		if info == nil {
			url = "https://www.hackthebox.com/api/v4/release_arena/active"
			resp = utils.HtbGet(url)
			info = utils.ParseJsonMessage(resp, "info")
			if info == nil {
				fmt.Println("No machine is active")
				os.Exit(1)
			}
			infomap := info.(map[string]interface{})
			fmt.Print("ID : ")
			fmt.Println(infomap["id"])
			fmt.Print("Name : ")
			fmt.Println(infomap["name"])
			fmt.Print("Plan : ")
			fmt.Println(infomap["type"])
			fmt.Print("Server : ")
			fmt.Println(infomap["lab_server"])
			fmt.Print("Status : Running")
		}
		infomap := info.(map[string]interface{})
		fmt.Print("ID : ")
		fmt.Println(infomap["id"])
		fmt.Print("Name : ")
		fmt.Println(infomap["name"])
		fmt.Print("Plan : ")
		fmt.Println(infomap["type"])
		fmt.Print("Server : ")
		fmt.Println(infomap["lab_server"])
		fmt.Println("Status : Running")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
