package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func anagrams(strs []string) (res []string) {
	m := make(map[string][]string)
	for _, s := range strs {
		key := makeKey(s)
		if _, ok := m[key]; !ok {
			m[key] = make([]string, 0)
			m[key] = append(m[key], s)
		} else {
			m[key] = append(m[key], s)
		}
	}
	for key, val := range m {
		if len(m[key]) > 1 {
			res = append(res, val...)
		}
	}
	return
}

func makeKey(s string) string {
	var buf bytes.Buffer
	m := make(map[rune]int)
	for _, c := range s {
		if _, ok := m[c]; ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		if _, ok := m[i]; ok {
			buf.WriteString(string(i))
			buf.WriteString(strconv.Itoa(m[i]))
		}
	}
	return buf.String()
}

func main() {
	// s := "abcabd"
	// fmt.Println(makeKey(s))
	strs := []string{"abc", "pcd", "cba", "ssd", "psd", "cdp"}
	res := anagrams(strs)
	fmt.Println(res)
	for _, v := range res {
		fmt.Println(v)
	}
}
