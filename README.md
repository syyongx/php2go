# php2go
Use Golang to implement PHP's common built-in functions. About 100+ functions have been implemented.

## Install
```
go get github.com/syyongx/php2go
```

## PHP Functions
### Date/Time Functions
```
time
strtotime
date
sleep
usleep
```
### String Functions
```
strpos
stripos
strrpos
strripos
str_replace
ucfirst
ucwords
substr
strrev
number_format
chunk_split
str_word_count
wordwrap
strlen
mb_strlen
str_repeat
strstr
str_shuffle
trim
ltrim
rtrim
explode
strtoupper
strtolower
chr
ord
nl2br
json_encode
json_decode
addslashes
stripslashes
quotemeta
htmlentities
html_entity_decode
md5
md5_file
sha1
sha1_file
crc32
```
### URL Functions
```
parse_url
url_encode
url_decode
ralurlencode
rawurldecode
base64_encode
base64_decode
```
### Array(Slice/Map) Functions
```
array_fill
array_flip
array_keys
array_values
array_merge
array_chunk
array_pad
array_slice
array_rand
array_column
array_pop
array_shift
array_unshift
array_diff
array_combine
array_reverse
implode
```
### Math Functions
```
rand
round
pi
max
min
decbin
bin2dec
hex2bin
bin2hex
dechex
hexdec
decoct
Octdec
base_convert
is_nan
```
### Filesystem Functions
```
file_exists
is_file
is_dir
filesize
file_put_contents
file_get_contents
unlink
delete
copy
is_readable
is_writeable
rename
touch
mkdir
getcwd
realpath
basename
chmod
chown
fclose
filemtime
fgetcsv
```
### Variable handling Functions
```
is_numeric
```
### Other Functions
```
echo
uniqid
exec
exit
die
Ternary(condition bool, trueVal, falseVal interface{}) interface{}
```

## LICENSE
[MIT](https://github.com/syyongx/php2go/blob/master/LICENSE)