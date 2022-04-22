package cmd

import (
	"fmt"
	"main/configs"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Julius",
	Long:  `All software has versions. This is Julius's`,
	Run: func(cmd *cobra.Command, args []string) {
		version := fmt.Sprintf("Julius v%s", configs.GetVersion())
		fmt.Println(version)
	},
}
