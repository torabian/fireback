//go:build linux || darwin

package backup

import "syscall"

// availableDiskSpace uses statfs(2), available on both linux and darwin via
// the stdlib syscall package (no golang.org/x/sys needed here) - Bavail is
// blocks available to an unprivileged user (what actually matters for "can
// this dump fit"), not Bfree (which includes blocks reserved for root).
// Bsize/Bavail's concrete field widths differ between the two OSes'
// syscall.Statfs_t, but both convert cleanly to uint64 regardless.
func availableDiskSpace(dir string) (int64, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(dir, &stat); err != nil {
		return 0, err
	}
	return int64(uint64(stat.Bavail) * uint64(stat.Bsize)), nil
}
