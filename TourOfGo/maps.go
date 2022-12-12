package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mydict := map[string]int{}
	mystr := strings.Fields(s)
	for i, _ := range mystr {
		_, ok := mydict[mystr[i]]
		if ok == false {
			mydict[mystr[i]] = 1
		} else {
			mydict[mystr[i]] += 1
		}
	}

	return mydict
}

func main() {
	wc.Test(WordCount)
}
