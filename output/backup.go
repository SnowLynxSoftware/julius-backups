package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"main/internal/util"
	"os"
)

func init() {
	rootCmd.AddCommand(backupCmd)
}

var backupCmd = &cobra.Command{
	Use:   "backup [source] [target]",
	Short: "Backup a source directory to the target directory.",
	Long: `Given a source directory, we will copy all of the files
to the target directory.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir := args[0]
		targetDir := args[1]

		targetFiles, err := ioutil.ReadDir(targetDir)
		if err != nil {
			panic(err)
		}

		if len(targetFiles) > 0 {
			fmt.Println("Target directory is not empty! Please choose a target directory that is empty and try again!")
			os.Exit(0)
		}

		sourceFiles, err := ioutil.ReadDir(sourceDir)
		if err != nil {
			panic(err)
		}

		for _, file := range sourceFiles {
			if !file.IsDir() {
				util.FileCopy(sourceDir, file, targetDir)
			}
		}
	},
}
