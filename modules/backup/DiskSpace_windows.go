//go:build windows

package backup

import "golang.org/x/sys/windows"

// availableDiskSpace uses GetDiskFreeSpaceEx - freeBytesAvailable (not
// totalFreeBytes) is the figure that respects the calling user's disk
// quota, matching what Bavail (vs Bfree) gives on unix in DiskSpace_unix.go.
func availableDiskSpace(dir string) (int64, error) {
	dirPtr, err := windows.UTF16PtrFromString(dir)
	if err != nil {
		return 0, err
	}

	var freeBytesAvailable, totalBytes, totalFreeBytes uint64
	if err := windows.GetDiskFreeSpaceEx(dirPtr, &freeBytesAvailable, &totalBytes, &totalFreeBytes); err != nil {
		return 0, err
	}
	return int64(freeBytesAvailable), nil
}
