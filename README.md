# PHP2Go

[![GoDoc](https://godoc.org/github.com/syyongx/php2go?status.svg)](https://godoc.org/github.com/syyongx/php2go)
[![Go Report Card](https://goreportcard.com/badge/github.com/syyongx/php2go)](https://goreportcard.com/report/github.com/syyongx/php2go)
[![MIT licensed][3]][4]

[3]: https://img.shields.io/badge/license-MIT-blue.svg
[4]: LICENSE

Use Golang to implement PHP's common built-in functions. About 140+ functions have been implemented.

## Install
```shell
go get github.com/syyongx/php2go
```

## Requirements
Go 1.10 or above.

## PHP Functions

### Date/Time Functions
```php
time()
strtotime()
date()
checkdate()
sleep()
usleep()
```

### String Functions
```php
strpos()
stripos()
strrpos()
strripos()
str_replace()
ucfirst()
lcfirst()
ucwords()
substr()
strrev()
number_format()
chunk_split()
str_word_count()
wordwrap()
strlen()
mb_strlen()
str_repeat()
strstr()
strtr()
str_shuffle()
trim()
ltrim()
rtrim()
explode()
strtoupper()
strtolower()
chr()
ord()
nl2br()
json_encode()
json_decode()
addslashes()
stripslashes()
quotemeta()
htmlentities()
html_entity_decode()
md5()
md5_file()
sha1()
sha1_file()
crc32()
levenshtein()
similar_text()
soundex()
parse_str()
```

### URL Functions
```php
base64_encode()
base64_decode()
parse_url()
urlencode()
urldecode()
rawurlencode()
rawurldecode()
http_build_query()
```

### Array(Slice/Map) Functions
```php
array_fill()
array_flip()
array_keys()
array_values()
array_merge()
array_chunk()
array_pad()
array_slice()
array_rand()
array_column()
array_push()
array_pop()
array_unshift()
array_shift()
array_key_exists()
array_combine()
array_reverse()
implode()
in_array()
```

### Mathematical Functions
```php
abs()
rand()
round()
floor()
ceil()
pi()
max()
min()
decbin()
bindec()
hex2bin()
bin2hex()
dechex()
hexdec()
decoct()
Octdec()
base_convert()
is_nan()
```

### CSPRNG Functions
```php
random_bytes()
random_int()
```

### Directory/Filesystem Functions
```php
stat()
pathinfo()
file_exists()
is_file()
is_dir()
filesize()
file_put_contents()
file_get_contents()
unlink()
delete()
copy()
is_readable()
is_writeable()
rename()
touch()
mkdir()
getcwd()
realpath()
basename()
chmod()
chown()
fclose()
filemtime()
fgetcsv()
glob()
```

### Variable handling Functions
```php
empty()
is_numeric()
```

### Program execution Functions
```php
exec()
system()
passthru()
```

### Network Functions
```php
gethostname()
gethostbyname()
gethostbynamel()
gethostbyaddr()
ip2long()
long2ip()
```

### Misc. Functions
```php
echo()
uniqid()
exit()
die()
getenv()
putenv()
memory_get_usage()
memory_get_peak_usage()
version_compare()
zip_open()
Ternary(condition bool, trueVal, falseVal interface{}) interface{}
```

## LICENSE
PHP2Go source code is licensed under the [MIT](https://github.com/syyongx/php2go/blob/master/LICENSE) Licence.
