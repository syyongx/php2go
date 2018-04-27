# php2go
Use Golang to implement PHP's common built-in functions. About 120+ functions have been implemented.

Hope this project will help PHPer to learn Golang.

## Install
```shell
go get github.com/syyongx/php2go
```

## PHP Functions
### Date/Time Functions
```php
time()
strtotime()
date()
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
```
### URL Functions
```php
parse_url()
urlencode()
urldecode()
rawurlencode()
rawurldecode()
base64_encode()
base64_decode()
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
disk_free_space()
disk_total_space()
glob()
umask()
```
### Variable handling Functions
```php
is_numeric()
```
### Other Functions
```php
echo()
uniqid()
exec()
exit()
die()
getenv()
putenv()
memory_get_usage()
version_compare()
zip_open()
Ternary(condition bool, trueVal, falseVal interface{}) interface{}
```

## LICENSE
php2go source code is licensed under the [MIT](https://github.com/syyongx/php2go/blob/master/LICENSE) Licence.