package charmap

import "fmt"


var KnownCharsets = map[string](*DecodeTable) {
	"CP1250"		: &CP1250_UTF8_TABLE,
	"cp1250"		: &CP1250_UTF8_TABLE,
	"WINDOWS-1250"	: &CP1250_UTF8_TABLE,
	"windows-1250"	: &CP1250_UTF8_TABLE,

	"CP1251"		: &CP1251_UTF8_TABLE,
	"cp1251"		: &CP1251_UTF8_TABLE,
	"WINDOWS-1251"	: &CP1251_UTF8_TABLE,
	"windows-1251"	: &CP1251_UTF8_TABLE,

	"CP1252"		: &CP1252_UTF8_TABLE,
	"cp1252"		: &CP1252_UTF8_TABLE,
	"WINDOWS-1252"	: &CP1252_UTF8_TABLE,
	"windows-1252"	: &CP1252_UTF8_TABLE,

	"CP852"			: &CP852_UTF8_TABLE,
	"cp852"			: &CP852_UTF8_TABLE,
	"IBM852"		: &CP852_UTF8_TABLE,
	"ibm852"		: &CP852_UTF8_TABLE,

	"CP855"			: &CP855_UTF8_TABLE,
	"cp855"			: &CP855_UTF8_TABLE,
	"IBM855"		: &CP855_UTF8_TABLE,
	"ibm855"		: &CP855_UTF8_TABLE,

	"CP866"			: &CP866_UTF8_TABLE,
	"cp866"			: &CP866_UTF8_TABLE,
	"IBM866"		: &CP866_UTF8_TABLE,
	"ibm866"		: &CP866_UTF8_TABLE,

	"KOI8-R"		: &KOI8R_UTF8_TABLE,
	"KOI8R"			: &KOI8R_UTF8_TABLE,
	"koi8-r"		: &KOI8R_UTF8_TABLE,
	"koi8r"			: &KOI8R_UTF8_TABLE,

	"KOI8-U"		: &KOI8U_UTF8_TABLE,
	"KOI8U"			: &KOI8U_UTF8_TABLE,
	"koi8-u"		: &KOI8U_UTF8_TABLE,
	"koi8u"			: &KOI8U_UTF8_TABLE,

	"ISO-8859-5"	: &ISO88595_UTF8_TABLE,
	"ISO8859-5"		: &ISO88595_UTF8_TABLE,
	"iso-8859-5"	: &ISO88595_UTF8_TABLE,
	"iso8859-5"		: &ISO88595_UTF8_TABLE,
}


func ToUTF8(table *DecodeTable, src []byte) []byte {
	var (
		dst = make([]byte, len(src)*3)
		tab = *table
		pos = 0
	)

	for _, b := range src {
		code := tab[b]
		for code > 0 {
			dst[pos] = byte(code)
			code >>= 8
			pos++
		}
	}
	return dst[:pos]
}

func ANY_to_UTF8(src []byte, charset string) ([]byte, error) {
	if table, ok := KnownCharsets[charset]; ok {
		return ToUTF8(table, src), nil
	} else {
		return src, fmt.Errorf("unknown charset %v", charset)
	}
}
