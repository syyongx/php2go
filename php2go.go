// php2go functions
package php2go

import (
	"time"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/url"
	"strings"
	"math"
	"encoding/base64"
	"encoding/json"
	"os"
	"strconv"
	"os/exec"
	"syscall"
	"bytes"
	"unicode/utf8"
	"path/filepath"
	"math/rand"
	"crypto/sha1"
	"hash/crc32"
	"html"
	"unicode"
	"io"
	"encoding/csv"
)

//////////// Date/Time Functions ////////////

// time()
func Time() int64 {
	return time.Now().Unix()
}

// strtotime()
func Strtotime(format, strtime string) (int64, error) {
	t, err := time.Parse(format, strtime)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// date()
func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}

// sleep()
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// usleep()
func Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}

//////////// String Functions ////////////

// strpos()
func Strpos(haystack, needle string, offset int) int {
	return strings.Index(haystack, needle)
}

// stripos()
func Stripos(haystack, needle string, offset int) int {
	haystack = strings.ToLower(haystack)
	needle = strings.ToLower(needle)
	return strings.Index(haystack, needle)
}

// strrpos()
func Strrpos(haystack, needle string, offset int) int {
	return strings.LastIndex(haystack, needle)
}

// strripos()
func Strripos(haystack, needle string, offset int) int {
	haystack = strings.ToLower(haystack)
	needle = strings.ToLower(needle)
	return strings.LastIndex(haystack, needle)
}

// str_replace()
func StrReplace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

// ucfirst()
func Ucfirst(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)
	f := strings.ToUpper(string(runes[0]))
	if len(runes) > 1 {
		return f + string(runes[1:])
	}
	return f
}

// ucwords()
func Ucwords(str string) string {
	return strings.Title(str)
}

// substr()
func Substr(str string, start uint, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

// strrev()
func Strrev(str string) string {
	runes := []rune(str)
	l := utf8.RuneCountInString(str)
	ns := make([]rune, l)
	for i, _ := range runes {
		l--
		ns[i] = runes[l]
	}
	return string(ns)
}

// number_format()
func NumberFormat() {

}

// chunk_split()
func ChunkSplit(body string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(body), []rune(end)
	l := uint(len(runes))
	if l <= 1 || l < chunklen {
		return body + end
	}
	var ns []rune
	var i uint
	for i = 0; i < l; i += chunklen {
		if i+chunklen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}

// str_word_count()
func StrWordCount(str string) []string {
	return strings.Fields(str)
}

// wordwrap()
func Wordwrap(str string, width uint, br string) string {
	if br == "" {
		br = "\n"
	}
	init := make([]byte, 0, len(str))
	buf := bytes.NewBuffer(init)
	var current uint
	var wordbuf, spacebuf bytes.Buffer
	for _, char := range str {
		if char == '\n' {
			if wordbuf.Len() == 0 {
				if current+uint(spacebuf.Len()) > width {
					current = 0
				} else {
					current += uint(spacebuf.Len())
					spacebuf.WriteTo(buf)
				}
				spacebuf.Reset()
			} else {
				current += uint(spacebuf.Len() + wordbuf.Len())
				spacebuf.WriteTo(buf)
				spacebuf.Reset()
				wordbuf.WriteTo(buf)
				wordbuf.Reset()
			}
			buf.WriteRune(char)
			current = 0
		} else if unicode.IsSpace(char) {
			if spacebuf.Len() == 0 || wordbuf.Len() > 0 {
				current += uint(spacebuf.Len() + wordbuf.Len())
				spacebuf.WriteTo(buf)
				spacebuf.Reset()
				wordbuf.WriteTo(buf)
				wordbuf.Reset()
			}
			spacebuf.WriteRune(char)
		} else {
			wordbuf.WriteRune(char)
			if current+uint(spacebuf.Len()+wordbuf.Len()) > width && uint(wordbuf.Len()) < width {
				buf.WriteString(br)
				current = 0
				spacebuf.Reset()
			}
		}
	}

	if wordbuf.Len() == 0 {
		if current+uint(spacebuf.Len()) <= width {
			spacebuf.WriteTo(buf)
		}
	} else {
		spacebuf.WriteTo(buf)
		wordbuf.WriteTo(buf)
	}

	return buf.String()
}

// strlen()
func Strlen(str string) int {
	return len(str)
}

// mb_strlen()
func MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}

// str_repeat()
func StrRepeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}

// strstr()
func Strstr(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return ""
	}
	return haystack[idx+len([]byte(needle))-1: ]
}

// str_shuffle()
func StrShuffle(str string) string {
	runes := []rune(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]rune, len(runes))
	for i, v := range r.Perm(len(runes)) {
		s[i] = runes[v]
	}
	return string(s)
}

// trim()
func Trim(str, characterMask string) string {
	if characterMask == "" {
		characterMask = " \\t\\n\\r\\0\\x0B"
	}
	return string(bytes.Trim([]byte(str), characterMask))
}

// ltrim()
func Ltrim(str, characterMask string) string {
	if characterMask == "" {
		characterMask = " \\t\\n\\r\\0\\x0B"
	}
	return string(bytes.TrimLeft([]byte(str), characterMask))
}

// rtrim()
func Rtrim(str, characterMask string) string {
	if characterMask == "" {
		characterMask = " \\t\\n\\r\\0\\x0B"
	}
	return string(bytes.TrimRight([]byte(str), characterMask))
}

// explode()
func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

// strtoupper()
func Strtoupper(str string) string {
	return strings.ToUpper(str)
}

// strtolower()
func Strtolower(str string) string {
	return strings.ToLower(str)
}

// chr()
func Chr(ascii int) string {
	return string(ascii)
}

// ord()
func Ord(char string) int {
	r, _ := utf8.DecodeRune([]byte(char))
	return int(r)
}

// nl2br()
// \n\r, \r\n, \r, \n
func Nl2br(str string, isXhtml bool) string {
	r, n, runes := '\r', '\n', []rune(str)
	var br []byte
	if isXhtml {
		br = []byte("<br />")
	} else {
		br = []byte("<br>")
	}
	skip := false
	var buf bytes.Buffer
	for i, v := range runes {
		if skip {
			skip = false
			continue
		}
		switch v {
		case n, r:
			if (i+1 < len(runes)) && (v == r && runes[i+1] == n) || (v == n && runes[i+1] == r) {
				buf.Write(br)
				skip = true
				continue
			}
			buf.Write(br)
		default:
			buf.WriteRune(v)
		}
	}
	return buf.String()
}

// json_encode()
func JsonEncode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}

// json_decode()
func JsonDecode(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}

// addslashes()
func Addslashes(str string) string {
	var n []rune
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			n = append(n, '\\')
		}
		n = append(n, char)
	}
	return string(n)
}

// stripslashes()
func Stripslashes(str string) string {
	return ""
}

// quotemeta()
func Quotemeta(str string) string {
	var n []rune
	for _, char := range str {
		switch char {
		case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
			n = append(n, '\\')
		}
		n = append(n, char)
	}
	return string(n)
}

// htmlentities()
func Htmlentities(str string) string {
	return html.EscapeString(str)
}

// html_entity_decode()
func HtmlEntityDecode(str string) string {
	return html.UnescapeString(str)
}

// md5()
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// md5_file()
func Md5File(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// sha1()
func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// sha1_file()
func Sha1File(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// crc32()
func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

//////////// URL Functions ////////////

// parse_url()
func ParseUrl(str string) (*url.URL, error) {
	return url.Parse(str)
}

// url_encode()
func UrlEncode(str string) string {
	return strings.Replace(url.PathEscape(str), "%20", "+", -1)
}

// url_decode()
func UrlDecode(str string) (string, error) {
	return url.PathUnescape(strings.Replace(str, "+", "%20", -1))
}

// ralurlencode()
func Rawurlencode(str string) string {
	return url.PathEscape(str)
}

// rawurldecode()
func RawurlDecode(str string) (string, error) {
	return url.PathUnescape(str)
}

// base64_encode()
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// base64_decode()
func Base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//////////// Array(Slice/Map) Functions ////////////

// array_fill()
func ArrayFill(startIndex int, num uint, value interface{}) map[int]interface{} {
	m := make(map[int]interface{})
	var i uint
	for i = 0; i < num; i++ {
		m[startIndex] = value
		startIndex++
	}
	return m
}

// array_flip()
func ArrayFlip(m map[interface{}]interface{}) map[interface{}]interface{} {
	newmap := make(map[interface{}]interface{})
	for i, v := range m {
		newmap[v] = i
	}

	return newmap
}

// array_keys()
func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
	var keys []interface{}
	for key, _ := range elements {
		keys = append(keys, key)
	}
	return keys
}

// array_values()
func ArrayValues(elements map[interface{}]interface{}) []interface{} {
	var vals []interface{}
	for _, val := range elements {
		vals = append(vals, val)
	}
	return vals
}

// array_merge()
func ArrayMerge(ss ...[]interface{}) []interface{} {
	var s []interface{}
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

// array_chunk()
func ArrayChunk(s []interface{}, size int) [][]interface{} {
	if size < 1 {
		panic("size: cannot be less than 1")
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]interface{}
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i ++
	}
	return n
}

// array_pad()
func ArrayPad(s []interface{}, size int, val interface{}) []interface{} {
	if size == 0 || (size > 0 && size < len(s)) || (size < 0 && size > -len(s)) {
		return s
	}
	n := size
	if size < 0 {
		n = -size
	}
	tmp := make([]interface{}, n)
	for i := 0; i < n; i++ {
		tmp[i] = val
	}
	if size > 0 {
		return append(s, tmp...)
	} else {
		return append(tmp, s...)
	}
}

// array_slice()
func ArraySlice(s []interface{}, offset, length uint) []interface{} {
	if offset > uint(len(s)) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < uint(len(s)) {
		return s[offset:end]
	}
	return s[offset:]
}

// array_rand()
func ArrayRand(elements []interface{}) []interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := make([]interface{}, len(elements))
	for i, v := range r.Perm(len(elements)) {
		n[i] = elements[v]
	}
	return n
}

// array_column()
func ArrayColumn(input map[string]map[string]interface{}, columnKey string) []interface{} {
	var res []interface{}
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			res = append(res, v)
		}
	}
	return res
}

// array_pop()
// for slice, map is unordered
func ArrayPop(elements []interface{}) interface{} {
	return elements[0:len(elements)-1]
}

// array_shift()
// for slice, map is unordered
func ArrayShift(elements []interface{}) []interface{} {
	return elements[1:]
}

// array_unshift()
func ArrayUnshift() {

}

// array_diff()
func ArrayDiff(maps ... map[interface{}]interface{}) {

}

// array_combine()
func ArrayCombine(s1, s2 []interface{}) map[interface{}]interface{} {
	if len(s1) != len(s2) {
		panic("the number of elements for each slice isn't equal")
	}
	m := make(map[interface{}]interface{}, len(s1))
	for i, v := range s1 {
		m[v] = s2[i]
	}

	return m
}

// array_reverse()
func ArrayReverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// implode()
func Implode(glue string, pieces []string) string {
	var buf bytes.Buffer
	l := len(pieces)
	for _, str := range pieces {
		buf.WriteString(str)
		if l--; l > 0 {
			buf.WriteString(glue)
		}
	}

	return buf.String()
}

//////////// Math Functions ////////////

// rand()
func Rand(min, max int) int {
	if min > max {
		panic("min: min mast less than max")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(math.MaxInt32)
	return n/((math.MaxInt32+1)/(max-min+1)) + min
}

// round()
func Round(f float64) float64 {
	return math.Floor(f + 0.5)
}

// pi()
func Pi() float64 {
	return math.Pi
}

// max()
func Max(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		max = math.Max(max, nums[i])
	}
	return max
}

// min()
func Min(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		min = math.Min(min, nums[i])
	}
	return min
}

// decbin()
func Decbin(number int64) string {
	return strconv.FormatInt(number, 2)
}

// bin2dec()
func Bin2dec(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 10), nil
}

// hex2bin()
func Hex2bin(data string) (string, error) {
	i, err := strconv.ParseInt(data, 16, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}

// bin2hex()
func Bin2hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

// dechex()
func Dechex(number int64) string {
	return strconv.FormatInt(number, 16)
}

// hexdec()
func Hexdec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 0)
}

// decoct()
func Decoct(number int64) string {
	return strconv.FormatInt(number, 8)
}

// Octdec()
func Octdec(str string) (int64, error) {
	return strconv.ParseInt(str, 8, 0)
}

// base_convert()
func BaseConvert(number string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(number, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// is_nan()
func IsNan(val float64) bool {
	return math.IsNaN(val)
}

//////////// Filesystem Functions ////////////

// file_exists()
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// is_file()
func IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// is_dir()
func IsDir(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	fm := fd.Mode()
	return fm.IsDir(), nil
}

// filesize()
func FileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

// file_put_contents()
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}

// file_get_contents()
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}

// unlink()
func Unlink(filename string) error {
	return os.Remove(filename)
}

// delete()
func Delete(filename string) error {
	return os.Remove(filename)
}

// copy()
func Copy(source, dest string) (bool, error) {
	fd1, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer fd1.Close()
	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	defer fd2.Close()
	_, e := io.Copy(fd2, fd1)
	if e != nil {
		return false, e
	}
	return true, nil
}

// is_readable
func IsReadable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_RDONLY, 0)
	if err != nil {
		return false
	}
	return true
}

// is_writeable()
func IsWriteable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_WRONLY, 0)
	if err != nil {
		return false
	}
	return true
}

// rename()
func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// touch()
func Touch(filename string) (bool, error) {
	fd, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return false, err
	}
	fd.Close()
	return true, nil
}

// mkdir()
func Mkdir(filename string, mode os.FileMode) error {
	return os.Mkdir(filename, mode)
}

// getcwd()
func Getcwd() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}

// realpath()
func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

// basename()
func Basename(path string) string {
	return filepath.Base(path)
}

// chmod()
func Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// chown()
func Chown(filename string, uid, gid int) bool {
	return os.Chown(filename, uid, gid) == nil
}

// fclose()
func Fclose(handle *os.File) error {
	return handle.Close()
}

// filemtime()
func Filemtime(filename string) (int64, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fd.Close()
	fileinfo, err := fd.Stat()
	if err != nil {
		return 0, err
	}
	return fileinfo.ModTime().Unix(), nil
}

// fgetcsv()
func Fgetcsv(handle *os.File, length int, delimiter rune) ([][]string, error) {
	reader := csv.NewReader(handle)
	reader.Comma = delimiter
	// TODO length limit
	return reader.ReadAll()
}

//////////// Variable handling Functions ////////////

// is_numeric()
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64:
	case uint, uint8, uint16, uint32, uint64:
	case float32, float64:
	case complex64, complex128:
		return true
	case string:
		// TODO check string is numeric
		return false
	default:
		return false
	}
	return false
}

//////////// Other Functions ////////////

// echo
func Echo(args ...interface{}) {
	fmt.Print(args...)
}

// uniqid()
func Uniqid(prefix string) string {
	t := time.Now()
	sec := t.Unix()
	usec := t.UnixNano() % 0x100000
	return fmt.Sprintf("%s%08x%05x", prefix, sec, usec);
}

// exec()
func Exec(command string) error {
	return exec.Command(command).Run()
}

// exit()
func Exit(status int) {
	os.Exit(status)
}

// die()
func Die(status int) {
	os.Exit(status)
}

// Ternary expression
// max := Ternary(a > b, a, b).(int)
func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
