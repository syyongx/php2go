package php2go

import (
	"testing"
)

func TestNetwork(t *testing.T) {
	tGethostname, _ := Gethostname()
	gt(t, float64(len(tGethostname)), 0)

	equal(t, uint32(134744072), IP2long("8.8.8.8"))

	equal(t, "8.8.8.8", Long2ip(134744072))

	tGethostbyname, _ := Gethostbyname("localhost")
	equal(t, "127.0.0.1", tGethostbyname)

	tGethostbynamel, _ := Gethostbynamel("localhost")
	gt(t, float64(len(tGethostbynamel)), 0)

	tGethostbyaddr, _ := Gethostbyaddr("127.0.0.1")
	equal(t, "localhost", tGethostbyaddr)
}
