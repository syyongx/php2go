package php2go

import "testing"

func TestProgramExecution(t *testing.T) {
	var output []string
	var retVal int
	gt(t, float64(len(Exec("/bin/bash -c \"ls -a|grep php\"", &output, &retVal))), 0)
	equal(t, 0, retVal)
	equal(t, 0, retVal)
	gt(t, float64(len(System("ls -l", &retVal))), 0)
	Passthru("echo hello", &retVal)
	equal(t, 0, retVal)
}
