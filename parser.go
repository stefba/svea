package main

import (
	"bytes"
	"unicode/utf8"
	"fmt"
)

func makeLinksLine(bts []byte) []byte {
	b := bytes.Buffer{}

	for len(bts) > 0 {
		r, size := utf8.DecodeRune(bts)
		bts = bts[size:]

		if r == '[' {
			x := makeLink(&b, bts)
			if x != -1 {
				bts = bts[x:]
				continue
			}
		}

		b.WriteRune(r)
	}

	return b.Bytes()
}

func makeLink(b *bytes.Buffer, bts []byte) int {
	name := bytes.Buffer{}
	href := bytes.Buffer{}

	inHref := false

	i := 0
	for len(bts) > 0 {
		r, size := utf8.DecodeRune(bts)
		bts = bts[size:]
		i += size

		if r == ')' && inHref {
			b.WriteString(fmt.Sprintf(`<a href="%v">%v</a>`, href.String(), name.String()))
			return i + 1
		}

		if r == ']' {
			inHref = true
			if len(bts) > 0 {
				if bts[0] != '(' {
					return -1
				}
				bts = bts[1:]
				continue
			}
			return -1
		}

		if inHref {
			href.WriteRune(r)
			continue
		}

		name.WriteRune(r)
	}
	return -1
}
