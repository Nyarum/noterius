#!/usr/bin/env python
# vim: fileencoding=utf-8 et ts=4 sts=4 sw=4 tw=0

from sys import stdout

CHARSETS = [
    "CP852",
    "CP855",
    "CP866",
    "CP1250",
    "CP1251",
    "CP1252",
    "KOI8-R",
    "KOI8-U",
    "ISO-8859-5",
]

def escape(u):
    e = u.encode("unicode_escape")
    if e.startswith(r"\x") and len(e) == 4:
        e = e[:-2] + e[-2:].upper()
    return e

def out(u, f=stdout):
    f.write(u.encode('UTF-8'))

def gen_utf8_table(charset):
    name = charset.replace('-','').upper()
    out("var %s_UTF8_TABLE = DecodeTable{\n" % name)
    for x in xrange(256):
        uni_chr   = chr(x).decode(charset, 'replace')
        printable = 0x20 <= ord(uni_chr) < 0x7F or ord(uni_chr) >= 0xA0
        uni_repr  = uni_chr if printable else escape(uni_chr)
        utf8_chr  = uni_chr.encode("UTF8", 'replace')
        utf8_hex  = u"0x%s," % ''.join(u"%02X" % ord(c) for c in reversed(utf8_chr))
        out(u"\t0x%02X : %-9s  // '%s'\n" % (x, utf8_hex, uni_repr))
    out("}\n")

def gen_decode_func(charset):
    name = charset.replace('-','').upper()
    out("""func %(name)s_to_UTF8(src []byte) []byte {
\treturn ToUTF8(&%(name)s_UTF8_TABLE, src)
}\n""" % locals())

if __name__ == "__main__":
    out("package charmap\n\n")

    out("type DecodeTable [256]uint32\n\n")

    for charset in CHARSETS:
        gen_utf8_table(charset)

    out("\n")

    for charset in CHARSETS:
        gen_decode_func(charset)

    out("\n")

