package portage

import (
	"testing"

	"github.com/anchore/syft/syft/artifact"
	"github.com/anchore/syft/syft/file"
	"github.com/anchore/syft/syft/pkg"
	"github.com/anchore/syft/syft/pkg/cataloger/internal/pkgtest"
	"github.com/anchore/syft/syft/source"
)

func TestPortageCataloger(t *testing.T) {

	expectedPkgs := []pkg.Package{
		{
			Name:    "app-containers/skopeo",
			Version: "1.5.1",
			FoundBy: "portage-cataloger",
			PURL:    "pkg:ebuild/app-containers/skopeo@1.5.1",
			Locations: source.NewLocationSet(
				source.NewLocation("var/db/pkg/app-containers/skopeo-1.5.1/CONTENTS"),
				source.NewLocation("var/db/pkg/app-containers/skopeo-1.5.1/LICENSE"),
				source.NewLocation("var/db/pkg/app-containers/skopeo-1.5.1/SIZE"),
			),
			Licenses:     []string{"Apache-2.0", "BSD", "BSD-2", "CC-BY-SA-4.0", "ISC", "MIT"},
			Type:         pkg.PortagePkg,
			MetadataType: pkg.PortageMetadataType,
			Metadata: pkg.PortageMetadata{
				InstalledSize: 27937835,
				Files: []pkg.PortageFileRecord{
					{
						Path: "/usr/bin/skopeo",
						Digest: &file.Digest{
							Algorithm: "md5",
							Value:     "376c02bd3b22804df8fdfdc895e7dbfb",
						},
					},
					{
						Path: "/etc/containers/policy.json",
						Digest: &file.Digest{
							Algorithm: "md5",
							Value:     "c01eb6950f03419e09d4fc88cb42ff6f",
						},
					},
					{
						Path: "/etc/containers/registries.d/default.yaml",
						Digest: &file.Digest{
							Algorithm: "md5",
							Value:     "e6e66cd3c24623e0667f26542e0e08f6",
						},
					},
					{
						Path: "/var/lib/atomic/sigstore/.keep_app-containers_skopeo-0",
						Digest: &file.Digest{
							Algorithm: "md5",
							Value:     "d41d8cd98f00b204e9800998ecf8427e",
						},
					},
				},
			},
		},
	}

	// TODO: relationships are not under test yet
	var expectedRelationships []artifact.Relationship

	pkgtest.NewCatalogTester().
		FromDirectory(t, "test-fixtures/image-portage").
		Expects(expectedPkgs, expectedRelationships).
		TestCataloger(t, NewPortageCataloger())

}
