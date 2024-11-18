package smi

import "github.com/hoony9x-furiosa-ai/68b4a517-730a-4911-b068-18f7cbcd58a7/pkg/smi/binding"

// VersionInfo represents a version information.
type VersionInfo interface {
	// Major returns a major part of version.
	Major() uint32
	// Minor returns a major part of version.
	Minor() uint32
	// Patch returns a major part of version.
	Patch() uint32
	// Metadata returns a major part of version.
	Metadata() string
}

var _ VersionInfo = new(versionInfo)

type versionInfo struct {
	raw binding.FuriosaSmiVersion
}

func newVersionInfo(raw binding.FuriosaSmiVersion) VersionInfo {
	return &versionInfo{raw: raw}
}

func (v *versionInfo) Major() uint32 {
	return v.raw.Major
}

func (v *versionInfo) Minor() uint32 {
	return v.raw.Minor
}

func (v *versionInfo) Patch() uint32 {
	return v.raw.Patch
}

func (v *versionInfo) Metadata() string {
	return byteBufferToString(v.raw.Metadata[:])
}
