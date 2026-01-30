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

var _ = Describe("Hash Command", func() {
	var (
		tempDir  string
		testFile string
		output   *bytes.Buffer
	)

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "hash-test")
		Expect(err).NotTo(HaveOccurred())

		output = new(bytes.Buffer)
		cmd.RootCmd.SetOut(output)

		testFile = filepath.Join(tempDir, "test.txt")
		Expect(ioutil.WriteFile(testFile, []byte("hello"), 0644)).To(Succeed())
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
		cmd.RootCmd.SetOut(nil)
	})

	It("should verify sha256 by default", func() {
		cmd.RootCmd.SetArgs([]string{"hash", testFile})
		err := cmd.RootCmd.Execute()
		Expect(err).NotTo(HaveOccurred())
		Expect(output.String()).To(ContainSubstring("2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"))
	})

	It("should verify md5", func() {
		cmd.RootCmd.SetArgs([]string{"hash", testFile, "--algo", "md5"})
		err := cmd.RootCmd.Execute()
		Expect(err).NotTo(HaveOccurred())
		Expect(output.String()).To(ContainSubstring("5d41402abc4b2a76b9719d911017c592"))
	})

	It("should fail for unsupported algorithm", func() {
		cmd.RootCmd.SetArgs([]string{"hash", testFile, "--algo", "invalid"})
		err := cmd.RootCmd.Execute()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("unsupported algorithm"))
	})
})
