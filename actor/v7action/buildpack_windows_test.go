//go:build windows
// +build windows

package v7action_test

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/integration/helpers"
	"code.cloudfoundry.org/ykk"
)

var _ = Describe("Buildpack", func() {
	Describe("Zipit", func() {
		var (
			source string
			target string

			executeErr error
		)

		JustBeforeEach(func() {
			executeErr = Zipit(source, target, "testzip-")
		})

		When("the source directory exists", func() {
			var subDir string

			BeforeEach(func() {
				var err error

				source, err = os.MkdirTemp("", "zipit-source-")
				Expect(err).ToNot(HaveOccurred())

				os.WriteFile(filepath.Join(source, "file1"), []byte{}, 0666)
				os.WriteFile(filepath.Join(source, "file2"), []byte{}, 0666)
				subDir, err = os.MkdirTemp(source, "zipit-subdir-")
				Expect(err).ToNot(HaveOccurred())
				os.WriteFile(filepath.Join(subDir, "file3"), []byte{}, 0666)

				p := filepath.FromSlash(fmt.Sprintf("buildpack-%s.zip", helpers.RandomName()))
				target, err = filepath.Abs(p)
				Expect(err).ToNot(HaveOccurred())
			})

			AfterEach(func() {
				Expect(os.RemoveAll(source)).ToNot(HaveOccurred())
				Expect(os.RemoveAll(target)).ToNot(HaveOccurred())
			})

			It("ensures that regular files have the executable bit set for the owner", func() {
				Expect(executeErr).ToNot(HaveOccurred())
				zipFile, err := os.Open(target)
				Expect(err).ToNot(HaveOccurred())
				defer zipFile.Close()

				zipStat, err := zipFile.Stat()
				reader, err := ykk.NewReader(zipFile, zipStat.Size())
				Expect(err).ToNot(HaveOccurred())

				Expect(reader.File).To(HaveLen(4))
				Expect(reader.File[0].Name).To(Equal("file1"))
				Expect(reader.File[0].Mode()).To(Equal(os.FileMode(0766)))

				Expect(reader.File[1].Name).To(Equal("file2"))
				Expect(reader.File[1].Mode()).To(Equal(os.FileMode(0766)))

				dirName := fmt.Sprintf("%s/", filepath.Base(subDir))
				Expect(reader.File[2].Name).To(Equal(dirName))
				Expect(reader.File[2].Mode()).To(Equal(os.FileMode(020000000777)))

				Expect(reader.File[3].Name).To(Equal(filepath.ToSlash(filepath.Join(dirName, "file3"))))
				Expect(reader.File[3].Mode()).To(Equal(os.FileMode(0766)))
			})
		})
	})
})
