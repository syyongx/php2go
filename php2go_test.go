package php2go

import (
	"testing"
	"fmt"
	"time"
	"os"
	"strconv"
)

func TestTime(t *testing.T) {

}

func TestString(t *testing.T) {
	tAddslashes := Addslashes("f'oo b\"ar")
	fmt.Println(tAddslashes)
	tStripslashes := Stripslashes("f\\'oo b\\\"ar\\\\a\\\\\\\\\\\\")
	fmt.Println(tStripslashes)
	tLevenshtein := Levenshtein("a", "b", 1, 1, 1)
	fmt.Println(tLevenshtein)
	var percent float64
	tSimilarText := SimilarText("a", "b", &percent)
	fmt.Println(tSimilarText, percent)
	tSoundex := Soundex("Heilbronn")
	fmt.Println(tSoundex)
}

func TestArray(t *testing.T) {
	var s1 = make([]interface{}, 3)
	s1[0] = "a"
	s1[1] = "b"
	s1[2] = "c"
	tArrayChunk := ArrayChunk(s1, 2)
	fmt.Println(tArrayChunk)
	var m1 = make(map[interface{}]interface{}, 3)
	m1[1] = "a"
	m1["a"] = "b"
	m1[2.5] = 1
	tArrayKeyExists := ArrayKeyExists("a", m1)
	fmt.Println(tArrayKeyExists)
	tArrayUnshift := ArrayUnshift(&s1, "x", "y")
	fmt.Println(tArrayUnshift)
	fmt.Println(s1)
	tArrayPush := ArrayPush(&s1, "u", "v")
	fmt.Println(tArrayPush)
	fmt.Println(s1)
	tArrayPop := ArrayPop(&s1)
	fmt.Println(tArrayPop)
	fmt.Println(s1)
	tArrayShift := ArrayShift(&s1)
	fmt.Println(tArrayShift)
	fmt.Println(s1)
}

func TestUrl(t *testing.T) {

}

func TestMath(t *testing.T) {
	tNumberFormat := NumberFormat(1234567890.777, 2, ".", ",")
	fmt.Println(tNumberFormat)
	tIsNumeric := IsNumeric("-0xaF")
	fmt.Println(tIsNumeric)
}

func TestOther(t *testing.T) {
	tVersionCompare := VersionCompare("1.3-beta", "1.4Rc1", "<")
	fmt.Println(tVersionCompare)
}

func TestFile(t *testing.T) {
	tPathinfo := Pathinfo("/home/go/.", 0)
	fmt.Println(tPathinfo)
	tDiskFreeSpace, _ := DiskFreeSpace("/")
	fmt.Println(tDiskFreeSpace)
	tDiskTotalSpace, _ := DiskTotalSpace("/")
	fmt.Println(tDiskTotalSpace)
}

func TestAll(t *testing.T) {
	taddslashes := Addslashes("'wo'中\"chousha")
	fmt.Println(taddslashes)
	ttimestamp := Time()
	fmt.Println(ttimestamp)
	tmd5 := Md5("123456")
	fmt.Println(tmd5)
	tsha1 := Sha1("123456")
	fmt.Println(tsha1)
	tcrc32 := Crc32("123456")
	fmt.Println(tcrc32)
	tstrrepeat := StrRepeat("我是", 10)
	fmt.Println(tstrrepeat)
	tsubstr := Substr("abcdef", 0, 2)
	fmt.Println(tsubstr)
	tstrstr := Strstr("xxx@gmail.com", "@")
	fmt.Println(tstrstr)
	tucfirst := Ucfirst("kello world")
	fmt.Println(tucfirst)
	tstrpos := Strpos("hello world", "w", 0)
	fmt.Println(tstrpos)
	timplode := Implode(",", []string{"a", "b", "c"})
	fmt.Println(timplode)
	trealpath1, err1 := Realpath("/home/go/../go/test/../")
	trealpath2, err2 := Realpath("~/Documents/godev")
	fmt.Println(trealpath1, trealpath2, err1, err2)
	tbasename := Basename("/home/go/src/pkg/php2go.go")
	fmt.Println(tbasename)
	tarrayfill := ArrayFill(-3, 6, "aaa")
	fmt.Println(tarrayfill)
	fmt.Println()
	tarrayrand := ArrayRand([]interface{}{"a", "b", "c"})
	fmt.Println(tarrayrand)
	tuniqid := Uniqid("")
	fmt.Println(tuniqid)
	tdate := Date("02/01/2006 15:04:05 PM", time.Now().Unix())
	fmt.Println(tdate)
	tstrtotime, _ := Strtotime("02/01/2006 15:04:05", "02/01/2016 15:04:05")
	fmt.Println(tstrtotime)
	tstrshuffle := StrShuffle("我˚abc")
	fmt.Println(tstrshuffle)
	tord := Ord("我是")
	fmt.Println(tord)
	tchr := Chr(25105)
	fmt.Println(tchr)
	tmbstrlen := MbStrlen("我是 a")
	fmt.Println(tmbstrlen)
	wd, _ := os.Getwd()
	tfilesize, _ := FileSize(wd)
	fmt.Println("filesize:" + strconv.FormatInt(tfilesize, 10))
	tnl2br := Nl2br("<a>\n\rxxx\nyy\r我是\r\nn\r\nx", false)
	fmt.Println(tnl2br)
	var s1 = make([]interface{}, 5, 5)
	s1[0] = "a"
	s1[1] = "b"
	s1[2] = "c"
	tarraypad := ArrayPad(s1, -5, 2)
	fmt.Println(tarraypad)
	tArrayChunk := ArrayChunk(s1, 2)
	fmt.Println(tArrayChunk)
	var s2 = make([]interface{}, 5, 5)
	s2[0] = "x"
	tarraycombine := ArrayCombine(s1, s2)
	fmt.Println(tarraycombine)
	tstrrev := Strrev("abc \t nic %我是10.5()---")
	fmt.Println("Strrev", tstrrev)
	tchunksplit := ChunkSplit("abc \t nic %我是10.5()---", 3, "e")
	fmt.Println(tchunksplit)
	tquotemeta := Quotemeta(".+?[$](*)^我是")
	fmt.Println(tquotemeta)
	tHtmlentities := Htmlentities("<html>hello world </html>")
	fmt.Println(tHtmlentities)
	tHtmlEntityDecode := HtmlEntityDecode("&lt;html&gt;hello world &lt;/html&gt;")
	fmt.Println(tHtmlEntityDecode)
	tWordwrap := Wordwrap("abc hello world xxx", 5, "")
	fmt.Println(tWordwrap)
	tStrWordCount := StrWordCount("a b c")
	fmt.Println(tStrWordCount)
	tMax := Max(2, 3.7, 5, 1.1)
	fmt.Println(tMax)
	tMin := Min(2, 3.7, 5, 1.1)
	fmt.Println(tMin)
	tRand := Rand(2, 5)
	fmt.Println(tRand)
	tDecbin := Decbin(100)
	fmt.Println(tDecbin)
	tBindec, _ := Bindec(tDecbin)
	fmt.Println(tBindec)
	tBin2hex, _ := Bin2hex(tDecbin)
	fmt.Println(tBin2hex)
	tHexdec, _ := Hexdec(tBin2hex)
	fmt.Println(tHexdec)
	tHex2bin, _ := Hex2bin(tBin2hex)
	fmt.Println(tHex2bin)
	tDecoct := Decoct(tHexdec)
	fmt.Println(tDecoct)
	tOctdec, _ := Octdec(tDecoct)
	fmt.Println(tOctdec)
	tDechex := Dechex(tHexdec)
	fmt.Println(tDechex)
	tBaseConvert, _ := BaseConvert(tBin2hex, 2, 16)
	fmt.Println(tBaseConvert)
	tFilePutContents := FilePutContents("", "", 066)
	fmt.Println(tFilePutContents)
}
