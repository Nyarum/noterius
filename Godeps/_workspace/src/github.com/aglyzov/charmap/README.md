## Golang library for fast conversion between 8-bit charsets and UTF-8

This little library allows to convert text from various 8-bit encodings
to UTF-8. And to make it **fast**. The code's using translation tables
to convert straight from an 8-bit encoding to UTF-8 without recoding to
slice of runes or other redundant steps.

### Supported charsets

* CP852 (aka IBM852)
* CP855 (aka IBM855)
* CP866 (aka IBM866)
* CP1250 (aka windows-1250)
* CP1251 (aka windows-1251)
* CP1252 (aka windows-1252)
* ISO-8859-5
* KOI8-R
* KOI8-U

### Extensibility

The translation tables are **generated** with a Python script and thus the
list of supported charsets **can easily be extended**. Just add necessary
charset to the list in [gen_tables.py](https://github.com/aglyzov/charmap/blob/master/gen_tables.py)
and run: `./gen_tables.py > tables.go`

### Examples

```go
import "github.com/aglyzov/charmap"

var koi8_r = "\xf4\xc5\xd3\xd4\xcf\xd7\xc1\xd1 \xd3\xd4\xd2\xcf\xcb\xc1"
var utf8_1 = charmap.ANY_to_UTF8(koi8_r, "KOI8-R")

// utf8_1 == "Тестовая строка"

var cp1251 = "\xd2\xe5\xf1\xf2\xee\xe2\xe0\xff \xf1\xf2\xf0\xee\xea\xe0"
var utf8_2 = charmap.CP1251_to_UTF8(cp1251)

// utf8_2 == utf8_1

var ibm866 = "\x92\xa5\xe1\xe2\xae\xa2\xa0\xef \xe1\xe2\xe0\xae\xaa\xa0"
var utf8_3 = charmap.ToUTF8(charmap.CP866_UTF8_TABLE, ibm866)

// utf8_3 == utf8_2
```
