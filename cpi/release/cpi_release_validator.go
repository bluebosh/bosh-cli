package release

import (
	"fmt"

	biinstallmanifest "github.com/cloudfoundry/bosh-init/installation/manifest"
	bitarball "github.com/cloudfoundry/bosh-init/installation/tarball"
	birel "github.com/cloudfoundry/bosh-init/release"
	birelmanifest "github.com/cloudfoundry/bosh-init/release/manifest"
	birelsetmanifest "github.com/cloudfoundry/bosh-init/release/set/manifest"
	biui "github.com/cloudfoundry/bosh-init/ui"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type CPIReleaseValidator interface {
	DownloadAndRegister(cpiReleaseRef birelmanifest.ReleaseRef, installationManifest biinstallmanifest.Manifest, stage biui.Stage) error
}

type cpiReleaseValidator struct {
	tarballProvider  bitarball.Provider
	releaseExtractor birel.Extractor
	releaseManager   birel.Manager
}

func NewCPIReleaseValidator(
	tarballProvider bitarball.Provider,
	releaseExtractor birel.Extractor,
	releaseManager birel.Manager,
) CPIReleaseValidator {
	return &cpiReleaseValidator{
		tarballProvider:  tarballProvider,
		releaseExtractor: releaseExtractor,
		releaseManager:   releaseManager,
	}
}

func (c *cpiReleaseValidator) DownloadAndRegister(cpiReleaseRef birelmanifest.ReleaseRef, installationManifest biinstallmanifest.Manifest, stage biui.Stage) error {
	releasePath, err := c.tarballProvider.Get(cpiReleaseRef, stage)
	if err != nil {
		return err
	}

	return stage.Perform(fmt.Sprintf("Validating release '%s'", cpiReleaseRef.Name), func() error {
		cpiRelease, err := c.releaseExtractor.Extract(releasePath)
		if err != nil {
			return bosherr.WrapErrorf(err, "Extracting release '%s'", releasePath)
		}

		c.releaseManager.Add(cpiRelease)
		err = NewValidator().Validate(cpiRelease, installationManifest.Template.Name)
		if err != nil {
			return bosherr.WrapErrorf(err, "Invalid CPI release '%s'", cpiRelease.Name())
		}
		return nil
	})
}

type ValidatedCpiReleaseSpec interface {
	GetFrom(deploymentManifestPath string) (biinstallmanifest.Manifest, birelmanifest.ReleaseRef, error)
}

type validatedCpiReleaseSpec struct {
	releaseSetManifestParser      birelsetmanifest.Parser
	releaseSetManifestValidator   birelsetmanifest.Validator
	installationManifestParser    biinstallmanifest.Parser
	installationManifestValidator biinstallmanifest.Validator
}

func NewValidatedCpiReleaseSpec(
	releaseSetParser birelsetmanifest.Parser,
	releaseSetValidator birelsetmanifest.Validator,
	installationParser biinstallmanifest.Parser,
	installationValidator biinstallmanifest.Validator,
) ValidatedCpiReleaseSpec {
	return &validatedCpiReleaseSpec{
		releaseSetManifestParser:      releaseSetParser,
		releaseSetManifestValidator:   releaseSetValidator,
		installationManifestParser:    installationParser,
		installationManifestValidator: installationValidator,
	}
}

func (v *validatedCpiReleaseSpec) GetFrom(deploymentManifestPath string) (biinstallmanifest.Manifest, birelmanifest.ReleaseRef, error) {
	releaseSetManifest, err := birelsetmanifest.ParseAndValidateFrom(deploymentManifestPath, v.releaseSetManifestParser, v.releaseSetManifestValidator)
	if err != nil {
		return biinstallmanifest.Manifest{}, birelmanifest.ReleaseRef{}, err
	}

	installationManifest, err := biinstallmanifest.ParseAndValidateFrom(deploymentManifestPath, v.installationManifestParser, v.installationManifestValidator, releaseSetManifest)
	if err != nil {
		return biinstallmanifest.Manifest{}, birelmanifest.ReleaseRef{}, err
	}

	cpiReleaseName := installationManifest.Template.Release
	cpiReleaseRef, found := releaseSetManifest.FindByName(cpiReleaseName)
	if !found {
		return biinstallmanifest.Manifest{}, birelmanifest.ReleaseRef{}, bosherr.Errorf("installation release '%s' must refer to a release in releases", cpiReleaseName)
	}
	return installationManifest, cpiReleaseRef, nil

}
