package php2go

import (
	"os"
	"testing"
)

func TestRealpath(t *testing.T) {
	pwd, _ := os.Getwd()
	os.Chdir("C:/")
	tRealpath, _ := Realpath("/")
	equal(t, "C:\\", tRealpath)
	tRealpath, _ = Realpath("/windows/fonts/../system32/")
	equal(t, "C:\\windows\\system32", tRealpath)
	os.Chdir(pwd)
}
func TestPathinfo(t *testing.T) {
	tPathInfo := Pathinfo("/home/go/php2go.go.go", -1)
	equal(t, map[string]string{"dirname": "\\home\\go", "basename": "php2go.go.go",
		"extension": "go", "filename": "php2go.go"}, tPathInfo)
}
