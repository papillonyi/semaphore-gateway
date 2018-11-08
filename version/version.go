package version

import "github.com/coreos/go-semver/semver"

var (
	// VersionMajor is for an API incompatible changes.
	VersionMajor int64
	// VersionMinor is for functionality in a backwards-compatible manner.
	VersionMinor int64 = 8
	// VersionPatch is for backwards-compatible bug fixes.
	VersionPatch int64 = 6
	// VersionPre indicates prerelease.
	VersionPre string
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string
)

var Version = semver.Version{
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,
	PreRelease: semver.PreRelease(VersionPre),
	Metadata:   VersionDev,
}


