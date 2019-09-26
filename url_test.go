package php2go

import "testing"

func TestUrl(t *testing.T) {
	tParseURL, _ := ParseURL("http://username:password@hostname:9090/path?arg=value#anchor", -1)
	equal(t, map[string]string{"pass": "password", "path": "/path", "query": "arg=value", "fragment": "anchor", "scheme": "http", "host": "hostname", "port": "9090", "user": "username"}, tParseURL)

	equal(t, "http%3A%2F%2Fgolang.org%3Fx+y", URLEncode("http://golang.org?x y"))

	tURLDecode, _ := URLDecode("http%3A%2F%2Fgolang.org%3Fx+y")
	equal(t, "http://golang.org?x y", tURLDecode)

	equal(t, "http%3A%2F%2Fgolang.org%3Fx%20y", Rawurlencode("http://golang.org?x y"))

	tRawurldecode, _ := Rawurldecode("http%3A%2F%2Fgolang.org%3Fx%20y")
	equal(t, "http://golang.org?x y", tRawurldecode)

	equal(t, "VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw==", Base64Encode("This is an encoded string"))

	tBase64Decode, _ := Base64Decode("VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw")
	equal(t, "This is an encoded string", tBase64Decode)

	equal(t, "first=value&multi=foo+bar&multi=baz", HTTPBuildQuery(map[string][]string{"first": {"value"}, "multi": {"foo bar", "baz"}}))
}
