package cmd

import (
	"fmt"
	"os"

	"github.com/QU35T-code/htb-cli/utils"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Save the machine chosen as an argument",
	Run: func(cmd *cobra.Command, args []string) {
		machine_id := utils.SearchMachineIDByName(args[0])
		os.Setenv("HTB_MACHINE_ID", machine_id)
		fmt.Println("The machine is correctly configured")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
