package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [item]",
	Short: "Search for files and folders",
	Long:  `Search for files and folders matching a specific term.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		dir, _ := cmd.Flags().GetString("dir")
		searchTerm := args[0]
		recursive, _ := cmd.Flags().GetBool("recursive")
		fileType, _ := cmd.Flags().GetString("type")
		caseSensitive, _ := cmd.Flags().GetBool("case-sensitive")
		contentSearch, _ := cmd.Flags().GetBool("content")

		if !caseSensitive {
			searchTerm = strings.ToLower(searchTerm)
		}
		matches := 0
		searchFunc := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if fileType == "files" && info.IsDir() {
				if path == dir {
					return nil
				}
				return filepath.SkipDir
			}
			if fileType == "folders" && !info.IsDir() {
				return nil
			}

			if contentSearch {
				if info.IsDir() {
					if !recursive && path != dir {
						return filepath.SkipDir
					}
					return nil
				}

				file, err := os.Open(path)
				if err != nil {
					return nil
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				lineNum := 0
				foundInFile := false
				for scanner.Scan() {
					lineNum++
					line := scanner.Text()
					checkLine := line
					if !caseSensitive {
						checkLine = strings.ToLower(line)
					}

					if strings.Contains(checkLine, searchTerm) {
						relativePath, _ := filepath.Rel(dir, path)
						fmt.Fprintf(cmd.OutOrStdout(), "  [MATCH] %s:%d: %s\n", relativePath, lineNum, strings.TrimSpace(line))
						matches++
						foundInFile = true
					}
				}
				if foundInFile {
				}
				return nil
			}

			name := info.Name()
			if !caseSensitive {
				name = strings.ToLower(name)
			}

			if strings.Contains(name, searchTerm) {
				relativePath, _ := filepath.Rel(dir, path)
				if info.IsDir() {
					fmt.Fprintf(cmd.OutOrStdout(), "  [DIR] %s\n", relativePath)
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "  [FILE] %s (%d bytes)\n", relativePath, info.Size())
				}
				matches++
			}

			if !recursive && path != dir && info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		err := filepath.Walk(dir, searchFunc)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "Error during search: %v\n", err)
			return err
		}

		if matches == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No matches found")
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "Found %d matches\n", matches)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolP("recursive", "r", false, "Search recursively")
	searchCmd.Flags().StringP("type", "t", "all", "Type to search (all, files, folders)")
	searchCmd.Flags().BoolP("case-sensitive", "c", false, "Use case-sensitive search")
	searchCmd.Flags().Bool("content", false, "Search file content instead of filenames")
}
