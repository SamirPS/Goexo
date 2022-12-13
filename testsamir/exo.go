package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var Graph = make(map[string][]int)

func main() {
	Graph["0"] = []int{1, 2}
	Graph["1"] = []int{2, 3}
	Graph["2"] = []int{4}
	Graph["3"] = []int{4}
	Graph["4"] = []int{2, 5}
	Graph["5"] = []int{3}
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go breadthfirstSearch(Graph, i)

	}
	wg.Wait()

}

func contains(s []int, element int) bool {
	for _, v := range s {
		if v == element {
			return true
		}
	}

	return false
}

func breadthfirstSearch(Graph map[string][]int, s int) {
	var f []int
	var t []int

	f = append(f, s)

	for len(f) > 0 {
		e := f[len(f)-1]
		f = f[0 : len(f)-1]
		t = append(t, e)
		for _, v := range Graph[fmt.Sprint(e)] {
			if !(contains(f, v) || contains(t, v)) {
				f = append(f, v)
			}
		}
	}

	fmt.Printf("Here the result:%v if we start from %d\n", t, s)
	wg.Done()

}
