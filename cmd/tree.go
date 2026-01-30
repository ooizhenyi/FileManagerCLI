package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var TreeCmd = &cobra.Command{
	Use:   "tree [dir]",
	Short: "List contents of directories in a tree-like format",
	Long:  `Visualizes the directory structure in a tree-like format, similar to the standard tree command.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		info, err := os.Stat(dir)
		if err != nil {
			return fmt.Errorf("error accessing directory '%s': %v", dir, err)
		}
		if !info.IsDir() {
			return fmt.Errorf("'%s' is not a directory", dir)
		}

		depth, _ := cmd.Flags().GetInt("depth")
		dirsOnly, _ := cmd.Flags().GetBool("dirs-only")

		fmt.Fprintln(cmd.OutOrStdout(), dir)
		return printTree(cmd.OutOrStdout(), dir, "", 0, depth, dirsOnly)
	},
}

func init() {
	RootCmd.AddCommand(TreeCmd)

	TreeCmd.Flags().Int("depth", -1, "Limit the depth of the tree print generation")
	TreeCmd.Flags().Bool("dirs-only", false, "List directories only")
}

func printTree(out io.Writer, path string, prefix string, currentDepth int, maxDepth int, dirsOnly bool) error {
	if maxDepth != -1 && currentDepth >= maxDepth {
		return nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var items []os.DirEntry
	for _, file := range files {
		if dirsOnly && !file.IsDir() {
			continue
		}
		items = append(items, file)
	}

	for i, file := range items {
		isLast := i == len(items)-1
		connector := "├── "
		newPrefix := prefix + "│   "
		if isLast {
			connector = "└── "
			newPrefix = prefix + "    "
		}

		fmt.Fprintf(out, "%s%s%s\n", prefix, connector, file.Name())

		if file.IsDir() {
			err := printTree(out, filepath.Join(path, file.Name()), newPrefix, currentDepth+1, maxDepth, dirsOnly)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
