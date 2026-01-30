package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash [file]",
	Short: "Calculate file checksum",
	Long:  `Calculate and print the cryptographic hash (checksum) of a file. Supported algorithms: md5, sha1, sha256.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]
		algo, _ := cmd.Flags().GetString("algo")

		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file '%s': %v", filePath, err)
		}
		defer file.Close()

		var h hash.Hash
		switch algo {
		case "md5":
			h = md5.New()
		case "sha1":
			h = sha1.New()
		case "sha256":
			h = sha256.New()
		default:
			return fmt.Errorf("unsupported algorithm: %s. Use md5, sha1, or sha256", algo)
		}

		if _, err := io.Copy(h, file); err != nil {
			return fmt.Errorf("failed to calculate hash: %v", err)
		}

		checksum := hex.EncodeToString(h.Sum(nil))
		fmt.Fprintf(cmd.OutOrStdout(), "%s  %s\n", checksum, filePath)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(hashCmd)
	hashCmd.Flags().String("algo", "sha256", "Hash algorithm (md5, sha1, sha256)")
}
