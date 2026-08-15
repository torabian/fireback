package backup

import (
	"os"
	"path/filepath"
)

// AvailableDiskSpace returns the free space, in bytes, on the filesystem
// that contains path. path itself doesn't need to exist yet (the dump
// output file usually doesn't, before the dump runs) - nearestExistingDir
// walks up to whichever ancestor directory actually does, and that's what
// gets statted. The real per-OS work is in DiskSpace_unix.go/
// DiskSpace_windows.go since there's no portable stdlib call for free disk
// space - shown to the operator before an interactive `backup dump` runs,
// alongside Dump.go's EstimateDatabaseSize, so an obviously-too-small
// target is visible up front rather than discovered mid-dump.
func AvailableDiskSpace(path string) (int64, error) {
	dir, err := nearestExistingDir(path)
	if err != nil {
		return 0, err
	}
	return availableDiskSpace(dir)
}

func nearestExistingDir(path string) (string, error) {
	dir, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	for {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return dir, nil // reached the filesystem root - stat it as-is and let the OS call itself report any error
		}
		dir = parent
	}
}
