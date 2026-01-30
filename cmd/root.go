package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "filemanager",
	Short: "A file management CLI application",
	Long: `A file management CLI application that allows you to:
- Create, move, delete, and rename folders
- View folder properties and contents
- Search for files and folders
- Copy, move, and delete files
- Compress and extract folders
- And more upcoming!`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringP("dir", "d", ".", "Directory to work with")
}
