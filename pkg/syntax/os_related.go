package syntax

import "runtime"

// OsRelatedValue returns the value of the os related parameter.
func OsRelatedValue[T any](windows T, linux T, darwin T, others T) T {
	switch runtime.GOOS {
	case "windows":
		return windows
	case "linux":
		return linux
	case "darwin":
		return darwin
	default:
		return others
	}
}

// ArchRelatedValue returns the value of the os and arch related parameter.
func ArchRelatedValue[T any](x64 T, x86 T, arm64 T, arm T, others T) T {
	switch runtime.GOARCH {
	case "amd64":
		return x64
	case "386":
		return x86
	case "arm64":
		return arm64
	case "arm":
		return arm
	default:
		return others
	}
}
