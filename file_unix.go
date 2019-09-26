// php2go functions

// +build linux darwin

package php2go

import (
	"syscall"
)

// Umask umask()
func Umask(mask int) int {
	return syscall.Umask(mask)
}

// DiskFreeSpace disk_free_space()
func DiskFreeSpace(directory string) (uint64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(directory, &fs)
	if err != nil {
		return 0, err
	}
	return fs.Bfree * uint64(fs.Bsize), nil
}

// DiskTotalSpace disk_total_space()
func DiskTotalSpace(directory string) (uint64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(directory, &fs)
	if err != nil {
		return 0, err
	}
	return fs.Blocks * uint64(fs.Bsize), nil
}
