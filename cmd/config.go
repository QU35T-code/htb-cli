package cmd

import (
	"github.com/spf13/cobra"
	"htb/utils"
	"fmt"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Save the machine chosen as an argument",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.SearchMachineIDByName(args[0])
		utils.SetConfigValue("machineID", machine_id)
		fmt.Println("The machine is correctly configured")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
