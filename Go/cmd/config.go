package cmd

import (
	"github.com/spf13/cobra"
	"htb/utils"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.SearchMachineIDByName(args[0])
		utils.SetConfigValue("machineID", machine_id)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
