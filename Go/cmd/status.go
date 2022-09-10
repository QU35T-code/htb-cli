package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"htb/utils"
	"os"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {		
		url := "https://www.hackthebox.com/api/v4/machine/active"
		resp := utils.HtbGet(url)
		info := utils.ParseJsonMessage(resp, "info")
		if (info == nil) {
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
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
