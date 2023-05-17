package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var mp map[string]int
	mp = make(map[string]int)
	for _, w := range strings.Fields(s) {
		if _, ok := mp[w]; ok {
			mp[w]++
		} else {
			mp[w] = 1
		}
	}
	return mp
}

func main() {
	wc.Test(WordCount)
}
