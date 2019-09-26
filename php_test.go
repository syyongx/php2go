package php2go

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestVariable(t *testing.T) {
	equal(t, true, IsNumeric("-0xaF"))
	equal(t, true, IsNumeric("123456"))

	equal(t, true, Empty(nil))
	equal(t, true, Empty(false))
	equal(t, true, Empty(0))
	equal(t, true, Empty(""))
	equal(t, true, Empty(0.0))
	equal(t, true, Empty([]int{}))
	equal(t, true, Empty([0]int{}))
	equal(t, false, Empty([1]int{}))
	equal(t, true, Empty(map[int]int{}))
}

func TestMisc(t *testing.T) {
	equal(t, true, VersionCompare("1.3-beta", "1.4Rc1", "<"))

	gt(t, float64(MemoryGetUsage(true)), 0)

	tPack, _ := Pack(binary.LittleEndian, []byte("abc"))
	equal(t, "abc", tPack)
	tUnpack, _ := Unpack(binary.LittleEndian, tPack)
	if _, ok := tUnpack.([]byte); !ok {
		t.Error("Unpack Failed")
	}
}

func BenchmarkFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChunkSplit("abc", 2, "\r\n")
	}
}

// Expected to be equal.
func equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
	}
}

// Expected to be unequal.
func unequal(t *testing.T, expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
	}
}

// Expect a greater than b.
func gt(t *testing.T, a, b float64) {
	if a <= b {
		t.Errorf("Expected %v (type %v) > Got %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

// Expect a greater than or equal to b.
func gte(t *testing.T, a, b float64) {
	if a < b {
		t.Errorf("Expected %v (type %v) > Got %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

// Expected value needs to be within range.
func rangeValue(t *testing.T, min, max, actual float64) {
	if actual < min || actual > max {
		t.Errorf("Expected range of %v-%v (type %v) > Got %v (type %v)", min, max, reflect.TypeOf(min), actual, reflect.TypeOf(actual))
	}
}
