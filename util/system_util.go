package util

import (
	"os"
	"path/filepath"
	"runtime"
)

//Get the absolute path to the running directory
func GetAppPath() string {
	if path, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		return path
	}
	return os.Args[0]
}

//Determine whether the current system is a Windows system?
func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
