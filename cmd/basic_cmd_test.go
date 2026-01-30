package cmd_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ooizhenyi/GoLangCLI/cmd"
)

var _ = Describe("Basic Commands", func() {
	var (
		tempDir string
		output  *bytes.Buffer
	)

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "basic-test")
		Expect(err).NotTo(HaveOccurred())

		output = new(bytes.Buffer)
		cmd.RootCmd.SetOut(output)
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
		cmd.RootCmd.SetOut(nil)
	})

	Context("Create Command", func() {
		It("should create a directory", func() {
			newDir := filepath.Join(tempDir, "newfolder")
			cmd.RootCmd.SetArgs([]string{"cf", "newfolder", "--dir", tempDir})
			err := cmd.RootCmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			info, err := os.Stat(newDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(info.IsDir()).To(BeTrue())
		})
	})

	Context("Delete Command", func() {
		It("should delete a directory", func() {
			dirToDelete := filepath.Join(tempDir, "deleteme")
			Expect(os.Mkdir(dirToDelete, 0755)).To(Succeed())

			cmd.RootCmd.SetArgs([]string{"dlt", "deleteme", "--dir", tempDir})
			err := cmd.RootCmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			_, err = os.Stat(dirToDelete)
			Expect(os.IsNotExist(err)).To(BeTrue())
		})
	})

	Context("Copy Command", func() {
		It("should copy a file", func() {
			srcFile := filepath.Join(tempDir, "source.txt")
			Expect(ioutil.WriteFile(srcFile, []byte("data"), 0644)).To(Succeed())

			cmd.RootCmd.SetArgs([]string{"copyfile", "source.txt", "dest.txt", "--dir", tempDir})
			err := cmd.RootCmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			destFile := filepath.Join(tempDir, "dest.txt")
			content, err := ioutil.ReadFile(destFile)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(Equal("data"))
		})
	})

	Context("Move Command", func() {
		It("should move a folder", func() {
			srcDir := filepath.Join(tempDir, "moveme")
			Expect(os.Mkdir(srcDir, 0755)).To(Succeed())

			cmd.RootCmd.SetArgs([]string{"mv", "moveme", "moved", "--dir", tempDir})
			err := cmd.RootCmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			_, err = os.Stat(srcDir)
			Expect(os.IsNotExist(err)).To(BeTrue())

			destDir := filepath.Join(tempDir, "moved")
			info, err := os.Stat(destDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(info.IsDir()).To(BeTrue())
		})
	})
})
