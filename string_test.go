package php2go

import (
	"bytes"
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	equal(t, `\'wo\'中文\"emoji`, Addslashes("'wo'中文\"emoji"))

	equal(t, "e10adc3949ba59abbe56e057f20f883e", Md5("123456"))

	equal(t, "7c4a8d09ca3762af61e59520943dc26494f8941b", Sha1("123456"))

	equal(t, uint32(158520161), Crc32("123456"))

	equal(t, "中文中文中文", StrRepeat("中文", 3))

	equal(t, "ab", Substr("abc", 0, 2))

	equal(t, "@gmail.com", Strstr("xxx@gmail.com", "@"))

	equal(t, "Hello world", Ucfirst("hello world"))

	equal(t, "hello world", Lcfirst("Hello world"))

	equal(t, "Hello World", Ucwords("hello world"))

	tStrlen := Strlen("G中文")
	equal(t, 7, tStrlen)

	tMbStrlen := MbStrlen("G中文")
	equal(t, 3, tMbStrlen)

	equal(t, 6, Strpos("hello wworld", "w", -6))

	equal(t, -1, Stripos("hello Wworld", "w", 8))

	equal(t, 6, Strrpos("hello wworld", "w", -6))

	equal(t, 7, Strripos("hello wWorld", "w", 0))

	equal(t, "a,b,c", Implode(",", []string{"a", "b", "c"}))

	tAddslashes := Addslashes("f'oo b\"ar")
	equal(t, `f\'oo b\"ar`, tAddslashes)

	equal(t, `f'oo b"ar\a\\\`, Stripslashes("f\\'oo b\\\"ar\\\\a\\\\\\\\\\\\"))

	equal(t, 4, Levenshtein("golang", "google", 1, 1, 1))

	var percent float64
	equal(t, 3, SimilarText("golang", "google", &percent))
	equal(t, float64(50), percent)

	equal(t, "H416", Soundex("Heilbronn"))

	equal(t, 14, len(Uniqid("x")))
	equal(t, true, bytes.HasPrefix([]byte(Uniqid("x")), []byte("x")))

	equal(t, 5, utf8.RuneCountInString(StrShuffle("中˚abc")))

	equal(t, 20013, Ord("中文"))

	equal(t, "z", Chr(122))

	equal(t, 4, MbStrlen("中文 a"))

	equal(t, "<a><br />xxx<br />yy<br />中文<br />n<br />x", Nl2br("<a>\n\rxxx\nyy\r中文\r\nn\r\nx", true))

	equal(t, "---)(5.01文中% cin 	 cba", Strrev("abc \t nic %中文10.5()---"))

	equal(t, "abce 	 enice %中e文10e.5(e)--e-e", ChunkSplit("abc \t nic %中文10.5()---", 3, "e"))

	equal(t, `\.\+\?\[\$\]\(\*\)\^中文`, Quotemeta(".+?[$](*)^中文"))

	equal(t, `&lt;html&gt;hello world &lt;/html&gt;`, Htmlentities("<html>hello world </html>"))

	equal(t, "<html>hello world </html>", HTMLEntityDecode("&lt;html&gt;hello world &lt;/html&gt;"))

	equal(t, "abc\nhello\nworld\nxxx", Wordwrap("abc hello world xxx", 5, "\n", false))

	equal(t, []string{"a", "b", "c"}, StrWordCount("a b c"))

	equal(t, "1001", Strtr("baab", "ab", "01"))
	equal(t, "bccb", Strtr("baab", "ab", "c"))
	equal(t, "bccb", Strtr("baab", "a", "cd"))
	equal(t, "ba01", Strtr("baab", map[string]string{"ab": "01"}))

	tParseStr := make(map[string]interface{})
	_ = ParseStr("f[a][]=m&f[a][]=n", tParseStr)
	equal(t, "map[f:map[a:[m n]]]", fmt.Sprint(tParseStr))
}
