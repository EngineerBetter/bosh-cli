package cmd_test

import (
	. "github.com/cloudfoundry/bosh-cli/cmd"

	boshrel "github.com/cloudfoundry/bosh-cli/release"
	boshjob "github.com/cloudfoundry/bosh-cli/release/job"
	boshpkg "github.com/cloudfoundry/bosh-cli/release/pkg"
	fakerel "github.com/cloudfoundry/bosh-cli/release/releasefakes"
	. "github.com/cloudfoundry/bosh-cli/release/resource"

	fakecrypto "github.com/cloudfoundry/bosh-cli/crypto/fakes"
	fakeui "github.com/cloudfoundry/bosh-cli/ui/fakes"
	fakefu "github.com/cloudfoundry/bosh-utils/fileutil/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/bosh-cli/crypto/fakes"
	"github.com/cloudfoundry/bosh-cli/release/license"
	"github.com/cloudfoundry/bosh-utils/errors"
)

var _ = FDescribe("Sha2ifyRelease", func() {

	var (
		releaseReader                *fakerel.FakeReader
		ui                           *fakeui.FakeUI
		fmv                          *fakefu.FakeMover
		releaseWriter                *fakerel.FakeWriter
		command                      Sha2ifyReleaseCmd
		args                         Sha2ifyReleaseArgs
		fakeDigestCalculator         *fakes.FakeDigestCalculator
		releaseWriterTempDestination string
	)

	BeforeEach(func() {
		releaseReader = &fakerel.FakeReader{}
		releaseWriter = &fakerel.FakeWriter{}
		ui = &fakeui.FakeUI{}
		fmv = &fakefu.FakeMover{}

		fakeDigestCalculator = fakes.NewFakeDigestCalculator()
		command = NewSha2ifyReleaseCmd(releaseReader, releaseWriter, fakeDigestCalculator, fmv, ui)
	})

	Context("Given a valid sha128 release tar", func() {
		var fakeSha128Release *fakerel.FakeRelease

		BeforeEach(func() {
			args = Sha2ifyReleaseArgs{
				Path:        "/some/release_128.tgz",
				Destination: "/some/release_256.tgz",
			}
			fakeSha128Release = &fakerel.FakeRelease{}
			jobSha128 := boshjob.NewJob(NewResourceWithBuiltArchive("job-resource-1", "job-sha128-fp", "/job-resource-1-path", "job-sha128-sha"))
			packageSha128 := boshpkg.NewPackage(NewResourceWithBuiltArchive("pkg-resource-1", "pkg-sha128-fp", "/pkg-resource-1-path", "pkg-sha128-sha"), nil)
			compiledPackageSha128 := boshpkg.NewCompiledPackageWithArchive("compiledpkg-resource-1", "compiledpkg-sha128-fp", "1", "/compiled-pkg-resource-path", "compiledpkg-sha128-sha", nil)

			fakeSha128Release.JobsReturns([]*boshjob.Job{jobSha128})
			fakeSha128Release.PackagesReturns([]*boshpkg.Package{packageSha128})
			fakeSha128Release.LicenseReturns(license.NewLicense(NewResourceWithBuiltArchive("license-resource-path", "lic-sha128-fp", "/license-resource-path", "lic-sha128-sha")))
			fakeSha128Release.CompiledPackagesReturns([]*boshpkg.CompiledPackage{compiledPackageSha128})

			fakeSha128Release.CopyWithStub = func(jobs []*boshjob.Job, pkgs []*boshpkg.Package, lic *license.License, compiledPackages []*boshpkg.CompiledPackage) boshrel.Release {
				fakeSha256Release := &fakerel.FakeRelease{}
				fakeSha256Release.JobsReturns(jobs)
				fakeSha256Release.PackagesReturns(pkgs)
				fakeSha256Release.LicenseReturns(lic)
				fakeSha256Release.CompiledPackagesReturns(compiledPackages)
				return fakeSha256Release
			}

			fakeDigestCalculator.SetCalculateBehavior(map[string]fakecrypto.CalculateInput{
				"/job-resource-1-path":        {DigestStr: "sha256:jobsha256"},
				"/pkg-resource-1-path":        {DigestStr: "sha256:pkgsha256"},
				"/license-resource-path":      {DigestStr: "sha256:licsha256"},
				"/compiled-pkg-resource-path": {DigestStr: "sha256:compiledpkgsha256"},
			})

			releaseReader.ReadReturns(fakeSha128Release, nil)
			releaseWriterTempDestination = "/some/temp/release_256.tgz"
			releaseWriter.WriteReturns(releaseWriterTempDestination, nil)

			err := command.Run(args)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Should convert it to a sha256 release tar", func() {
			Expect(releaseReader.ReadCallCount()).ToNot(Equal(0))

			readPathArg := releaseReader.ReadArgsForCall(0)
			Expect(readPathArg).To(Equal("/some/release_128.tgz"))

			Expect(releaseWriter.WriteCallCount()).To(Equal(1))
			sha2ifyRelease, _ := releaseWriter.WriteArgsForCall(0)

			Expect(sha2ifyRelease).NotTo(BeNil())

			Expect(sha2ifyRelease.License()).ToNot(BeNil())
			Expect(sha2ifyRelease.License().ArchiveSHA1()).To(Equal("sha256:licsha256"))

			Expect(sha2ifyRelease.Jobs()).To(HaveLen(1))
			Expect(sha2ifyRelease.Jobs()[0].ArchiveSHA1()).To(Equal("sha256:jobsha256"))

			Expect(sha2ifyRelease.Packages()).To(HaveLen(1))
			Expect(sha2ifyRelease.Packages()[0].ArchiveSHA1()).To(Equal("sha256:pkgsha256"))

			Expect(sha2ifyRelease.CompiledPackages()).To(HaveLen(1))
			Expect(sha2ifyRelease.CompiledPackages()[0].ArchiveSHA1()).To(Equal("sha256:compiledpkgsha256"))

			Expect(fmv.MoveCallCount()).To(Equal(1))

			src, dst := fmv.MoveArgsForCall(0)
			Expect(src).To(Equal(releaseWriterTempDestination))
			Expect(dst).To(Equal(args.Destination))

		})

		Context("when unable to write the sha256 tarball", func() {
			BeforeEach(func() {
				releaseWriter.WriteReturns("", errors.Error("disaster"))
			})

			It("should return an error", func() {
				err := command.Run(args)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("disaster"))
			})
		})

		Context("when rehashing a licence fails", func() {
			BeforeEach(func() {
				fakeDigestCalculator.SetCalculateBehavior(map[string]fakecrypto.CalculateInput{
					"/job-resource-1-path":        {DigestStr: "sha256:jobsha256"},
					"/pkg-resource-1-path":        {DigestStr: "sha256:pkgsha256"},
					"/compiled-pkg-resource-path": {DigestStr: "sha256:compiledpkgsha256"},
					"/license-resource-path":      {Err: errors.Error("Unknown algorithm")},
				})
			})

			It("should return an error", func() {
				err := command.Run(args)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Unknown algorithm"))
			})
		})

		Context("when rehashing compiled packages fails", func() {
			BeforeEach(func() {
				fakeDigestCalculator.SetCalculateBehavior(map[string]fakecrypto.CalculateInput{
					"/job-resource-1-path":        {DigestStr: "sha256:jobsha256"},
					"/pkg-resource-1-path":        {DigestStr: "sha256:pkgsha256"},
					"/compiled-pkg-resource-path": {Err: errors.Error("Unknown algorithm")},
					"/license-resource-path":      {DigestStr: "sha256:licsha256"},
				})
			})

			It("should return an error", func() {
				err := command.Run(args)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Unknown algorithm"))
			})
		})

		Context("when no licence is provided", func() {
			BeforeEach(func() {
				fakeSha128Release.LicenseReturns(nil)
				fakeDigestCalculator.SetCalculateBehavior(map[string]fakecrypto.CalculateInput{
					"/job-resource-1-path":        {DigestStr: "sha256:jobsha256"},
					"/pkg-resource-1-path":        {DigestStr: "sha256:pkgsha256"},
					"/compiled-pkg-resource-path": {DigestStr: "sha256:compiledpkgsha256"},
				})
			})

			It("should not return an error", func() {
				err := command.Run(args)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("When unable to move sha2fyied release to destination", func() {
			BeforeEach(func() {
				fmv.MoveReturns(errors.Error("disaster"))
			})

			It("Should return an error", func() {
				err := command.Run(args)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("disaster"))
			})
		})
	})

	Context("Given a bad file path", func() {
		BeforeEach(func() {
			args = Sha2ifyReleaseArgs{
				Path:        "/some/release_128.tgz",
				Destination: "/some/release_256.tgz",
			}

			releaseReader.ReadReturns(nil, errors.Error("disaster"))
		})

		It("Should return an error", func() {
			err := command.Run(args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("disaster"))
		})
	})
})
