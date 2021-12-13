/*
Copyright 2019 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/coreos/pkg/capnslog"
	"github.com/pkg/errors"
)

// CephVersion represents the Ceph version format
type CephVersion struct {
	Major int
	Minor int
	Extra int
	Build int
	IsNes bool
}

const (
	unknownVersionString = "<unknown version>"
)

var (
	// Minimum supported version is 14.2.5
	Minimum = CephVersion{14, 2, 5, 0, false}
	// Nautilus Ceph version
	Nautilus = CephVersion{14, 0, 0, 0, false}
	// Octopus Ceph version
	Octopus = CephVersion{15, 0, 0, 0, false}
	// Pacific Ceph version
	Pacific = CephVersion{16, 0, 0, 0, false}
	// Quincy Ceph version
	Quincy = CephVersion{17, 0, 0, 0, false}

	// Minimum supported version of NES is 1.0.0
	MinimumNES = CephVersion{1, 0, 0, 0, true}
	// Apex Ceph version
	Apex = CephVersion{1, 0, 0, 0, true}

	// cephVolumeLVMDiskSortingCephVersion introduced a major regression in c-v and thus is not suitable for production
	cephVolumeLVMDiskSortingCephVersion = CephVersion{Major: 14, Minor: 2, Extra: 13}

	// supportedVersions are production-ready versions that rook supports
	supportedVersions = []CephVersion{Nautilus, Apex, Octopus}

	// unsupportedVersions are possibly Ceph pin-point release that introduced breaking changes and not recommended
	unsupportedVersions = []CephVersion{cephVolumeLVMDiskSortingCephVersion}

	// for parsing the output of `ceph --version`
	versionPattern = regexp.MustCompile(`ceph version (\d+)\.(\d+)\.(\d+)(\.nes){0,1}`)

	// For a build release the output is "ceph version 14.2.4-64.el8cp"
	// So we need to detect the build version change
	buildVersionPattern = regexp.MustCompile(`ceph version (\d+)\.(\d+)\.(\d+)(\.nes){0,1}-(\d+)`)

	logger = capnslog.NewPackageLogger("github.com/rook/rook", "cephver")
)

func (v *CephVersion) String() string {
	var ret_v string

	ret_v = fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Extra)

	if v.IsNes {
		ret_v += ".nes"
	}

	ret_v += fmt.Sprintf("-%d %s", v.Build, v.ReleaseName())
	return ret_v
}

// CephVersionFormatted returns the Ceph version in a human readable format
func (v *CephVersion) CephVersionFormatted() string {
	return "ceph version " + v.String()
}

// ReleaseName is the name of the Ceph release
func (v *CephVersion) ReleaseName() string {
	if v.IsNes {
		switch v.Major {
		case Apex.Major:
			return "apex"
		default:
			return unknownVersionString
		}
	} else {
		switch v.Major {
		case Octopus.Major:
			return "octopus"
		case Nautilus.Major:
			return "nautilus"
		case Pacific.Major:
			return "pacific"
		case Quincy.Major:
			return "quincy"
		default:
			return unknownVersionString
		}
	}
}

// ExtractCephVersion extracts the major, minor and extra digit of a Ceph release
func ExtractCephVersion(src string) (*CephVersion, error) {
	var build int
	m := versionPattern.FindStringSubmatch(src)
	if m == nil {
		return nil, errors.Errorf("failed to parse version from: %q", src)
	}

	major, err := strconv.Atoi(m[1])
	if err != nil {
		return nil, errors.Errorf("failed to parse version major part: %q", m[1])
	}

	minor, err := strconv.Atoi(m[2])
	if err != nil {
		return nil, errors.Errorf("failed to parse version minor part: %q", m[2])
	}

	extra, err := strconv.Atoi(m[3])
	if err != nil {
		return nil, errors.Errorf("failed to parse version extra part: %q", m[3])
	}

	isnes := (m[4] == ".nes")

	// See if we are running on a build release
	mm := buildVersionPattern.FindStringSubmatch(src)
	// We don't need to handle any error here, so let's jump in only when "mm" has content
	if mm != nil {
		build, err = strconv.Atoi(mm[5])
		if err != nil {
			logger.Warningf("failed to convert version build number part %q to an integer, ignoring", mm[5])
		}
	}

	return &CephVersion{major, minor, extra, build, isnes}, nil
}

// Supported checks if a given release is supported
func (v *CephVersion) Supported() bool {
	for _, sv := range supportedVersions {
		if v.isRelease(sv) {
			return true
		}
	}
	return false
}

// Unsupported checks if a given release is supported
func (v *CephVersion) Unsupported() bool {
	for _, sv := range unsupportedVersions {
		if v.isExactly(sv) {
			return true
		}
	}
	return false
}

func (v *CephVersion) isRelease(other CephVersion) bool {
	return v.Major == other.Major && v.IsNes == other.IsNes
}

func (v *CephVersion) isExactly(other CephVersion) bool {
	return v.Major == other.Major && v.Minor == other.Minor && v.Extra == other.Extra && v.IsNes == other.IsNes
}

// IsNautilus checks if the Ceph version is Nautilus
func (v *CephVersion) IsNautilus() bool {
	return v.isRelease(Nautilus)
}

// IsOctopus checks if the Ceph version is Octopus
func (v *CephVersion) IsOctopus() bool {
	return v.isRelease(Octopus)
}

// IsPacific checks if the Ceph version is Pacific
func (v *CephVersion) IsPacific() bool {
	return v.isRelease(Pacific)
}

// IsQuincy checks if the Ceph version is Quincy
func (v *CephVersion) IsQuincy() bool {
	return v.isRelease(Quincy)
}

// IsApex checks if the NES version is Apex
func (v *CephVersion) IsApex() bool {
	return v.isRelease(Apex)
}

// IsAtLeast checks a given Ceph version is at least a given one
func (v *CephVersion) IsAtLeast(other CephVersion) bool {
	if v.Major > other.Major {
		return true
	} else if v.Major < other.Major {
		return false
	}
	// If we arrive here then v.Major == other.Major
	if v.Minor > other.Minor {
		return true
	} else if v.Minor < other.Minor {
		return false
	}
	// If we arrive here then v.Minor == other.Minor
	if v.Extra > other.Extra {
		return true
	} else if v.Extra < other.Extra {
		return false
	}
	// If we arrive here then both versions are identical
	return true
}

// IsAtLeastQuincy check that the Ceph version is at least Quincy
func (v *CephVersion) IsAtLeastQuincy() bool {
	return v.IsAtLeast(Quincy)
}

// IsAtLeastPacific check that the Ceph version is at least Pacific
func (v *CephVersion) IsAtLeastPacific() bool {
	return v.IsAtLeast(Pacific)
}

// IsAtLeastOctopus check that the Ceph version is at least Octopus
func (v *CephVersion) IsAtLeastOctopus() bool {
	return v.IsAtLeast(Octopus)
}

// IsAtLeastNautilus check that the Ceph version is at least Nautilus
func (v *CephVersion) IsAtLeastNautilus() bool {
	return v.IsAtLeast(Nautilus)
}

// IsAtLeastApex check that the NES version is at least Apex
func (v *CephVersion) IsAtLeastApex() bool {
	return v.IsAtLeast(Apex)
}

// IsIdentical checks if Ceph versions are identical
func IsIdentical(a, b CephVersion) bool {
	if a.Major == b.Major {
		if a.Minor == b.Minor {
			if a.Extra == b.Extra {
				if a.Build == b.Build {
					if a.IsNes == b.IsNes {
						return true
					}
				}
			}
		}
	}

	return false
}

// IsSuperior checks if a given version if superior to another one
func IsSuperior(a, b CephVersion) bool {
	if a.Major > b.Major {
		if a.IsNes == b.IsNes {
			return true
		}
	}
	if a.Major == b.Major {
		if a.Minor > b.Minor {
			if a.IsNes == b.IsNes {
				return true
			}
		}
	}
	if a.Major == b.Major {
		if a.Minor == b.Minor {
			if a.Extra > b.Extra {
				if a.IsNes == b.IsNes {
					return true
				}
			}
		}
	}
	if a.Major == b.Major {
		if a.Minor == b.Minor {
			if a.Extra == b.Extra {
				if a.Build > b.Build {
					if a.IsNes == b.IsNes {
						return true
					}
				}
			}
		}
	}

	return false
}

// IsInferior checks if a given version if inferior to another one
func IsInferior(a, b CephVersion) bool {
	if a.Major < b.Major {
		if a.IsNes == b.IsNes {
			return true
		}
	}
	if a.Major == b.Major {
		if a.Minor < b.Minor {
			if a.IsNes == b.IsNes {
				return true
			}
		}
	}
	if a.Major == b.Major {
		if a.Minor == b.Minor {
			if a.Extra < b.Extra {
				if a.IsNes == b.IsNes {
					return true
				}
			}
		}
	}
	if a.Major == b.Major {
		if a.Minor == b.Minor {
			if a.Extra == b.Extra {
				if a.Build < b.Build {
					if a.IsNes == b.IsNes {
						return true
					}
				}
			}
		}
	}

	return false
}

// ValidateCephVersionsBetweenLocalAndExternalClusters makes sure an external cluster can be connected
// by checking the external ceph versions available and comparing it with the local image provided
func ValidateCephVersionsBetweenLocalAndExternalClusters(localVersion, externalVersion CephVersion) error {
	logger.Debugf("local version is %q, external version is %q", localVersion.String(), externalVersion.String())

	// We only support Nautilus, Apex or newer
	if (externalVersion.IsNes && !externalVersion.IsAtLeastApex()) || (!externalVersion.IsNes && !externalVersion.IsAtLeastNautilus()) {
		return errors.Errorf("unsupported ceph version %q, need at least apex or nautilus, delete your cluster CR and create a new one with a correct ceph version", externalVersion.String())
	}

	// Identical version, regardless if other CRs are running, it's ok!
	if IsIdentical(localVersion, externalVersion) {
		return nil
	}

	// Local version must never be higher than the external one
	if IsSuperior(localVersion, externalVersion) {
		return errors.Errorf("local cluster ceph version is higher %q than the external cluster %q, this must never happen", externalVersion.String(), localVersion.String())
	}

	// External cluster was updated to a minor version higher, consider updating too!
	if localVersion.Major == externalVersion.Major {
		if IsSuperior(externalVersion, localVersion) {
			logger.Warningf("external cluster ceph version is a minor version higher %q than the local cluster %q, consider upgrading", externalVersion.String(), localVersion.String())
			return nil
		}
	}

	// The external cluster was upgraded, consider upgrading too!
	if localVersion.Major < externalVersion.Major {
		logger.Errorf("external cluster ceph version is a major version higher %q than the local cluster %q, consider upgrading", externalVersion.String(), localVersion.String())
		return nil
	}

	return nil
}
