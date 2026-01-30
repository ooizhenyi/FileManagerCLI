package cmd_test

import (
	"bytes"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ooizhenyi/GoLangCLI/cmd"
)

var _ = Describe("Tree Command", func() {
	var (
		tempDir string
		output  *bytes.Buffer
	)

	BeforeEach(func() {
		var err error
		tempDir, err = os.MkdirTemp("", "tree-test")
		Expect(err).NotTo(HaveOccurred())

		output = new(bytes.Buffer)
		cmd.RootCmd.SetOut(output)

		cmd.TreeCmd.Flags().Set("depth", "-1")
		cmd.TreeCmd.Flags().Set("dirs-only", "false")

		Expect(os.MkdirAll(filepath.Join(tempDir, "dir1", "subdir1"), 0755)).To(Succeed())
		Expect(os.WriteFile(filepath.Join(tempDir, "dir1", "file1.txt"), []byte("content"), 0644)).To(Succeed())
		Expect(os.WriteFile(filepath.Join(tempDir, "dir1", "subdir1", "file2.txt"), []byte("content"), 0644)).To(Succeed())
		Expect(os.WriteFile(filepath.Join(tempDir, "file3.txt"), []byte("content"), 0644)).To(Succeed())
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
		cmd.RootCmd.SetOut(nil)
	})

	It("should print the full tree structure", func() {
		cmd.RootCmd.SetArgs([]string{"tree", tempDir})
		err := cmd.RootCmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		out := output.String()
		Expect(out).To(ContainSubstring("dir1"))
		Expect(out).To(ContainSubstring("file1.txt"))
		Expect(out).To(ContainSubstring("subdir1"))
		Expect(out).To(ContainSubstring("file2.txt"))
		Expect(out).To(ContainSubstring("file3.txt"))
	})

	It("should respect depth flag", func() {
		cmd.RootCmd.SetArgs([]string{"tree", tempDir, "--depth", "1"})
		err := cmd.RootCmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		out := output.String()
		Expect(out).To(ContainSubstring("dir1"))
		Expect(out).To(ContainSubstring("file3.txt"))
	})

	It("should respect dirs-only flag", func() {
		cmd.RootCmd.SetArgs([]string{"tree", tempDir, "--dirs-only"})
		err := cmd.RootCmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		out := output.String()
		Expect(out).To(ContainSubstring("dir1"))
		Expect(out).To(ContainSubstring("subdir1"))
		Expect(out).NotTo(ContainSubstring("file1.txt"))
		Expect(out).NotTo(ContainSubstring("file3.txt"))
	})
})
