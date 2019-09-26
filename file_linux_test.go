package php2go

import (
	"testing"
)

func TestRealpath(t *testing.T) {
	tRealpath, _ := Realpath("/etc/fonts/../../proc/")
	equal(t, "/proc", tRealpath)
}
