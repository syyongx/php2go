package php2go

import "testing"

func TestMath(t *testing.T) {
	equal(t, float64(5), Max(2, 3.7, 5, 1.1))

	equal(t, 1.1, Min(2, 3.7, 5, 1.1))

	rangeValue(t, float64(2), float64(5), float64(Rand(2, 5)))

	tDecbin := Decbin(100)
	equal(t, "1100100", tDecbin)

	tBindec, _ := Bindec(tDecbin)
	equal(t, "100", tBindec)

	tBin2hex, _ := Bin2hex(tDecbin)
	equal(t, "64", tBin2hex)

	tHexdec, _ := Hexdec(tBin2hex)
	equal(t, int64(100), tHexdec)

	tHex2bin, _ := Hex2bin(tBin2hex)
	equal(t, "1100100", tHex2bin)

	tDecoct := Decoct(tHexdec)
	equal(t, "144", tDecoct)

	tOctdec, _ := Octdec(tDecoct)
	equal(t, int64(100), tOctdec)

	tDechex := Dechex(tHexdec)
	equal(t, "64", tDechex)

	tBaseConvert, _ := BaseConvert("64", 16, 2)
	equal(t, "1100100", tBaseConvert)

	equal(t, "1,234,567,890.78", NumberFormat(1234567890.777, 2, ".", ","))
}
