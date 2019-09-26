package php2go

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	equal(t, "php2go.go", Basename("/home/go/src/pkg/php2go.go"))
	wd, _ := os.Getwd()
	tFileSize, _ := FileSize(wd)
	gt(t, float64(tFileSize), 0)
}
