package system

import "runtime"

type Type int

const (
	Unknown Type = iota
	Windows
	Linux
	Darwin
)

// Detect os
func DetectOS() Type {
	switch runtime.GOOS {
	case "windows":
		return Windows
	case "linux":
		return Linux
	case "darwin":
		return Darwin
	default:
		return Unknown
	}
}
