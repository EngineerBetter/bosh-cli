package license_test

import (
	"errors"

	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-init/release/license"
	. "github.com/cloudfoundry/bosh-init/release/resource"
	fakeres "github.com/cloudfoundry/bosh-init/release/resource/fakes"
)

var _ = Describe("DirReaderImpl", func() {
	var (
		collectedFiles     []File
		collectedPrepFiles []File
		collectedChunks    []string
		archive            *fakeres.FakeArchive
		fs                 *fakesys.FakeFileSystem
		reader             DirReader
	)

	BeforeEach(func() {
		archive = &fakeres.FakeArchive{}
		archiveFactory := func(files, prepFiles []File, chunks []string) Archive {
			collectedFiles = files
			collectedPrepFiles = prepFiles
			collectedChunks = chunks
			return archive
		}
		fs = fakesys.NewFakeFileSystem()
		reader = NewDirReaderImpl(archiveFactory, fs)
	})

	Describe("Read", func() {
		It("returns a license collected from directory", func() {
			fs.WriteFileString("LICENSE", "license-content")

			fs.SetGlob("/dir/LICENSE*", []string{"/dir/LICENSE"})
			fs.SetGlob("/dir/NOTICE*", []string{})

			archive.FingerprintReturns("fp", nil)

			license, err := reader.Read("/dir")
			Expect(err).NotTo(HaveOccurred())
			Expect(license).To(Equal(NewLicense(NewResource("license", "fp", archive))))

			Expect(collectedFiles).To(Equal([]File{
				File{Path: "/dir/LICENSE", DirPath: "/dir", RelativePath: "LICENSE", UseBasename: true, ExcludeMode: true},
			}))

			Expect(collectedPrepFiles).To(BeEmpty())
			Expect(collectedChunks).To(BeEmpty())
		})

		It("returns a license and notice collected from directory", func() {
			fs.WriteFileString("LICENSE", "license-content")
			fs.WriteFileString("NOTICE", "notice-content")

			fs.SetGlob("/dir/LICENSE*", []string{"/dir/LICENSE"})
			fs.SetGlob("/dir/NOTICE*", []string{"/dir/NOTICE.md"})

			archive.FingerprintReturns("fp", nil)

			license, err := reader.Read("/dir")
			Expect(err).NotTo(HaveOccurred())
			Expect(license).To(Equal(NewLicense(NewResource("license", "fp", archive))))

			Expect(collectedFiles).To(Equal([]File{
				File{Path: "/dir/LICENSE", DirPath: "/dir", RelativePath: "LICENSE", UseBasename: true, ExcludeMode: true},
				File{Path: "/dir/NOTICE.md", DirPath: "/dir", RelativePath: "NOTICE.md", UseBasename: true, ExcludeMode: true},
			}))

			Expect(collectedPrepFiles).To(BeEmpty())
			Expect(collectedChunks).To(BeEmpty())
		})

		It("returns nil if there are no collected files", func() {
			license, err := reader.Read("/dir")
			Expect(err).NotTo(HaveOccurred())
			Expect(license).To(BeNil())
		})

		It("returns error if globbing fails", func() {
			fs.GlobErr = errors.New("fake-err")

			_, err := reader.Read("/dir")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("returns error if fingerprinting fails", func() {
			fs.WriteFileString("LICENSE", "license-content")
			fs.SetGlob("/dir/LICENSE*", []string{"/dir/LICENSE"})

			archive.FingerprintReturns("", errors.New("fake-err"))

			_, err := reader.Read("/dir")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
