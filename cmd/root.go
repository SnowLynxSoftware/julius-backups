package cmd

// CONFIG: https://github.com/spf13/cobra/blob/master/user_guide.md

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "julius",
	Short: "Julius is a FOSS backup utility.",
	Long: `Julius is a free and open source backup utility for UNIX 
systems built with love by MorpheusZero!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("For help, type: julius help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
