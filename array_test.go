package php2go

import "testing"

func TestArray(t *testing.T) {
	var s1 = make([]interface{}, 3)
	s1[0] = "a"
	s1[1] = "b"
	s1[2] = "c"
	tArrayChunk := ArrayChunk(s1, 2)
	equal(t, [][]interface{}{{"a", "b"}, {"c"}}, tArrayChunk)

	var m1 = make(map[interface{}]interface{}, 3)
	m1[1] = "a"
	m1["a"] = "b"
	m1[2.5] = 1
	tArrayKeyExists := ArrayKeyExists("a", m1)
	equal(t, true, tArrayKeyExists)

	tArrayUnshift := ArrayUnshift(&s1, "x", "y")
	equal(t, 5, tArrayUnshift)
	equal(t, []interface{}{"x", "y", "a", "b", "c"}, s1)

	equal(t, 7, ArrayPush(&s1, "u", "v"))
	equal(t, []interface{}{"x", "y", "a", "b", "c", "u", "v"}, s1)

	equal(t, "v", ArrayPop(&s1))
	equal(t, []interface{}{"x", "y", "a", "b", "c", "u"}, s1)

	tArrayShift := ArrayShift(&s1)
	equal(t, "x", tArrayShift)
	equal(t, []interface{}{"y", "a", "b", "c", "u"}, s1)

	equal(t, map[int]interface{}{-1: "aa", 0: "aa", 1: "aa", 2: "aa", -3: "aa", -2: "aa"}, ArrayFill(-3, 6, "aa"))

	equal(t, 3, len(ArrayRand([]interface{}{"a", "b", "c"})))

	var s2 = make([]interface{}, 3)
	s2[0] = "a"
	s2[1] = "b"
	s2[2] = "c"
	equal(t, []interface{}{"d", "d", "a", "b", "c"}, ArrayPad(s2, -5, "d"))

	var s3 = make([]interface{}, 3, 3)
	s3[0] = "x"
	s3[1] = "y"
	s3[2] = "z"
	equal(t, map[interface{}]interface{}{"a": "x", "b": "y", "c": "z"}, ArrayCombine(s2, s3))

	tInArray1 := InArray(1, [2]interface{}{"a", 1})                        // array
	tInArray2 := InArray(1, []interface{}{"a", 1})                         // slice
	tInArray3 := InArray(1, map[interface{}]interface{}{"a": "c", 1: "d"}) // map
	equal(t, true, tInArray1)
	equal(t, true, tInArray2)
	equal(t, false, tInArray3)
}
