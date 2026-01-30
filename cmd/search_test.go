package cmd_test

import (
	"bytes"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ooizhenyi/GoLangCLI/cmd"
)

var _ = Describe("Search Command", func() {
	var (
		tempDir string
		output  *bytes.Buffer
	)

	BeforeEach(func() {
		var err error
		tempDir, err = os.MkdirTemp("", "search-test")
		Expect(err).NotTo(HaveOccurred())

		output = new(bytes.Buffer)
		cmd.RootCmd.SetOut(output)

		Expect(os.WriteFile(filepath.Join(tempDir, "file1.txt"), []byte("Hello World"), 0644)).To(Succeed())
		Expect(os.WriteFile(filepath.Join(tempDir, "file2.txt"), []byte("Goodbye World"), 0644)).To(Succeed())
		Expect(os.WriteFile(filepath.Join(tempDir, "code.go"), []byte("func main() {}"), 0644)).To(Succeed())
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
		cmd.RootCmd.SetOut(nil)
	})

	Context("Filename Search", func() {
		It("should find files by name", func() {
			cmd.RootCmd.SetArgs([]string{"search", "file1", "--dir", tempDir})
			err := cmd.RootCmd.Execute()
			Expect(err).NotTo(HaveOccurred())
			Expect(output.String()).To(ContainSubstring("file1.txt"))
			Expect(output.String()).NotTo(ContainSubstring("file2.txt"))
		})
	})

	Context("Content Search", func() {
		It("should find files containing the search term", func() {
			cmd.RootCmd.SetArgs([]string{"search", "Goodbye", "--dir", tempDir, "--content", "--case-sensitive=false"})
			err := cmd.RootCmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			Expect(output.String()).To(ContainSubstring("file2.txt"))
			Expect(output.String()).NotTo(ContainSubstring("file1.txt"))
			Expect(output.String()).To(ContainSubstring("[MATCH]"))
		})

		It("should be case insensitive by default", func() {
			cmd.RootCmd.SetArgs([]string{"search", "goodbye", "--dir", tempDir, "--content"})
			err := cmd.RootCmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			Expect(output.String()).To(ContainSubstring("file2.txt"))
		})

		It("should respect case sensitivity", func() {
			cmd.RootCmd.SetArgs([]string{"search", "goodbye", "--dir", tempDir, "--content", "--case-sensitive=true"})
			err := cmd.RootCmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			Expect(output.String()).NotTo(ContainSubstring("file2.txt"))
			Expect(output.String()).To(ContainSubstring("No matches found"))
		})
	})
})
