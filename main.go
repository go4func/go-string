package main

import (
	"bytes"
	"log"
	"reflect"
	"sort"
	"strings"
)

func main() {
	var s string = "1234567.11"
	log.Println(comma(s))

	log.Println(anagram(s, "765.114321"))
}

// func comma(s string) string {
// 	n := len(s)
// 	if n <= 3 {
// 		return s
// 	}

// 	i := strings.LastIndex(s, ".")
// 	if i < 0 {
// 		return comma(s[:n-3]) + "," + s[n-3:]
// 	}

// 	s, f := s[:i], s[i:]
// 	return comma(s[:n-3]) + f
// }

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var f string
	if in := strings.LastIndex(s, "."); in >= 0 {
		s, f = s[:in], s[in:]
	}

	var b bytes.Buffer
	for i, v := range s {
		if (n-i)%3 == 0 {
			b.WriteString(",")
		}
		b.WriteRune(v)
	}

	return b.String() + f
}

func anagram(s, a string) bool {
	n := len(s)
	if len(a) != n {
		return false
	}

	ss := []rune(s)
	sa := []rune(a)
	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
	sort.Slice(sa, func(i, j int) bool { return sa[i] < sa[j] })

	return reflect.DeepEqual(ss, sa)
}
