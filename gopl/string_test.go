package gopl

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

// Compare

// EqualFold

// Contains

// Cont

// Fields

// Split Join
func Join(str []string, sep string) string {
	// 特殊情况应该做处理
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str[0]
	}
	buffer := bytes.NewBufferString(str[0])
	for _, s := range str[1:] {
		buffer.WriteString(sep)
		buffer.WriteString(s)
	}
	return buffer.String()
}
// HasPrefix  HasSuffix

// Repeat

//
func mmap(){
	mapping := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z': // 大写字母转小写
			return r + 32
		case r >= 'a' && r <= 'z': // 小写字母不处理
			return r
		case unicode.Is(unicode.Han, r): // 汉字换行
			return '\n'
		}
		return -1 // 过滤所有非字母、汉字的字符
	}
	fmt.Println(strings.Map(mapping, "Hello你#￥%……\n（'World\n,好Hello^(&(*界gopher..."))
}

//Replace  ReplaceAll

// ToLower ToUpper

// title

// Replacer

// Reader

// Builder

