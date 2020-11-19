package cheru

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var cheruSET string = "切卟叮咧哔唎啪啰啵嘭噜噼巴拉蹦铃"
var runecheruSET []rune = []rune(cheruSET)
var cheruMap map[rune]byte

func init() {
	cnt := byte(0)
	cheruMap = make(map[rune]byte)
	for _, char := range cheruSET {
		cheruMap[char] = cnt
		cnt++
	}
}

// 把单词转换成切噜词
func word2cheru(w string) string {
	res := strings.Builder{}
	res.WriteString("切")
	for _, char := range w {
		if unicode.IsPunct(char) {
			res.WriteRune(char)
			continue
		}
		runeW := []rune{}
		bytechars := []byte(string(char))
		for _, bytechar := range bytechars {
			runeW = append(runeW, runecheruSET[bytechar>>4&0xf])
			runeW = append(runeW, runecheruSET[bytechar&0xf])
		}
		for i := 0; i < len(runeW); i++ {
			res.WriteRune(runeW[i])
		}
	}
	return res.String()
}

// 把切噜词转换成单词
func cheru2word(w string) string {
	res := strings.Builder{}
	runeW := []rune(w)
	if runeW[0] != '切' || len(runeW) < 2 {
		return w
	}
	byteStr := []byte{}
	for i := 1; i < len(runeW)-1; i += 2 {
		x := cheruMap[runeW[i]] << 4
		x = x ^ cheruMap[runeW[i+1]]
		byteStr = append(byteStr, x)
	}
	for len(byteStr) != 0 {
		rune, size := utf8.DecodeRune(byteStr)
		res.WriteRune(rune)
		byteStr = byteStr[size:]
	}
	return res.String()
}

// Str2cheru 把字符串转换成切噜语
func Str2cheru(s string) string {
	res := strings.Builder{}
	words := afterFieldsFunc(s, unicode.IsPunct)
	for _, word := range words {
		res.WriteString(word2cheru(word))
	}
	return res.String()
}

// Cheru2str 把切噜语转换成字符串
func Cheru2str(s string) string {
	res := strings.Builder{}
	words := afterFieldsFunc(s, unicode.IsPunct)
	for _, word := range words {
		res.WriteString(cheru2word(word))
	}
	return res.String()
}

// 保留分隔符的切分
func afterFieldsFunc(s string, f func(rune) bool) []string {
	type span struct {
		start int
		end   int
	}
	spans := make([]span, 0, 32)

	start := -1 // valid span start if >= 0
	s2 := []rune(s)
	for end, rune := range s2 {
		if f(rune) {
			if start >= 0 {
				spans = append(spans, span{start, end})
				spans = append(spans, span{end, end + 1})
				start = ^start
			}
		} else {
			if start < 0 {
				start = end
			}
		}
	}

	// Last field might end at EOF.
	if start >= 0 {
		spans = append(spans, span{start, len(s2)})
	}

	// Create strings from recorded field indices.
	a := make([]string, len(spans))

	for i, span := range spans {
		a[i] = string(s2[span.start:span.end])
	}

	return a
}
