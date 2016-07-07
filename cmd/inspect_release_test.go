package cmd_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-init/cmd"
	boshdir "github.com/cloudfoundry/bosh-init/director"
	fakedir "github.com/cloudfoundry/bosh-init/director/fakes"
	fakeui "github.com/cloudfoundry/bosh-init/ui/fakes"
	boshtbl "github.com/cloudfoundry/bosh-init/ui/table"
)

var _ = Describe("InspectReleaseCmd", func() {
	var (
		ui       *fakeui.FakeUI
		director *fakedir.FakeDirector
		command  InspectReleaseCmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		director = &fakedir.FakeDirector{}
		command = NewInspectReleaseCmd(ui, director)
	})

	Describe("Run", func() {
		var (
			opts    InspectReleaseOpts
			release *fakedir.FakeRelease
		)

		BeforeEach(func() {
			opts = InspectReleaseOpts{
				Args: InspectReleaseArgs{
					Slug: boshdir.NewReleaseSlug("some-name", "some-version"),
				},
			}

			release = &fakedir.FakeRelease{}
			director.FindReleaseReturns(release, nil)
		})

		act := func() error { return command.Run(opts) }

		It("shows jobs and packages for specified release", func() {
			release.JobsStub = func() ([]boshdir.Job, error) {
				return []boshdir.Job{
					{
						Name:        "some-job-name",
						Fingerprint: "some-job-fingerprint",

						BlobstoreID: "some-job-blob-id",
						SHA1:        "some-job-blob-sha1",
					},
				}, nil
			}

			release.PackagesStub = func() ([]boshdir.Package, error) {
				return []boshdir.Package{
					{
						Name:        "some-pkg1-name",
						Fingerprint: "some-pkg1-fingerprint",

						BlobstoreID: "some-pkg1-blob-id",
						SHA1:        "some-pkg1-blob-sha1",
					},
					{
						Name:        "some-pkg2-name",
						Fingerprint: "some-pkg2-fingerprint",

						BlobstoreID: "some-pkg2-blob-id",
						SHA1:        "some-pkg2-blob-sha1",

						CompiledPackages: []boshdir.CompiledPackage{
							{
								StemcellSlug: boshdir.NewStemcellSlug(
									"some-stemcell-name",
									"some-stemcell-version",
								),

								BlobstoreID: "some-compiled-pkg-blob-id",
								SHA1:        "some-compiled-pkg-blob-sha1",
							},
						},
					},
				}, nil
			}

			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(director.FindReleaseArgsForCall(0)).To(Equal(
				boshdir.NewReleaseSlug("some-name", "some-version")))

			Expect(ui.Tables).To(Equal([]boshtbl.Table{
				{
					Content: "jobs",

					Header: []string{"Job", "Blobstore ID", "SHA1"},

					SortBy: []boshtbl.ColumnSort{{Column: 0, Asc: true}},

					Rows: [][]boshtbl.Value{
						{
							boshtbl.NewValueString("some-job-name/some-job-fingerprint"),
							boshtbl.NewValueString("some-job-blob-id"),
							boshtbl.NewValueString("some-job-blob-sha1"),
						},
					},
				},
				{
					Content: "packages",

					Header: []string{"Package", "Compiled for", "Blobstore ID", "SHA1"},

					SortBy: []boshtbl.ColumnSort{{Column: 0, Asc: true}},

					Sections: []boshtbl.Section{
						{
							FirstColumn: boshtbl.NewValueString("some-pkg1-name/some-pkg1-fingerprint"),

							Rows: [][]boshtbl.Value{
								{
									boshtbl.NewValueString(""),
									boshtbl.NewValueString("(source)"),
									boshtbl.NewValueString("some-pkg1-blob-id"),
									boshtbl.NewValueString("some-pkg1-blob-sha1"),
								},
							},
						},
						{
							FirstColumn: boshtbl.NewValueString("some-pkg2-name/some-pkg2-fingerprint"),

							Rows: [][]boshtbl.Value{
								{
									boshtbl.NewValueString(""),
									boshtbl.NewValueString("(source)"),
									boshtbl.NewValueString("some-pkg2-blob-id"),
									boshtbl.NewValueString("some-pkg2-blob-sha1"),
								},
								{
									boshtbl.NewValueString(""),
									boshtbl.NewValueString("some-stemcell-name/some-stemcell-version"),
									boshtbl.NewValueString("some-compiled-pkg-blob-id"),
									boshtbl.NewValueString("some-compiled-pkg-blob-sha1"),
								},
							},
						},
					},
				},
			}))
		})

		It("returns error if jobs cannot be retrieved", func() {
			release.JobsReturns(nil, errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("returns error if jobs cannot be retrieved", func() {
			release.PackagesReturns(nil, errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("returns error if release cannot be retrieved", func() {
			director.FindReleaseReturns(nil, errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})