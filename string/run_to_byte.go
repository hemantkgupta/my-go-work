package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

func main() {
	s := "Color Infection is a fun game."
	bs := []byte(s) // string -> []byte
	s = string(bs)  // []byte -> string

	rs := []rune(s) // string -> []rune
	s = string(rs)  // []rune -> string

	rs = bytes.Runes(bs) // []byte -> []rune
	fmt.Println(bs)
	fmt.Println(rs)

	bs = Runes2Bytes(rs) // []rune -> []byte
	fmt.Println(bs)
}
