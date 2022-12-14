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
	var tovisit []int
	var visited []int
	tovisit = append(tovisit, s)

	for len(tovisit) > 0 {
		node := tovisit[0]
		tovisit = tovisit[1:]
		visited = append(visited, node)
		for _, v := range Graph[fmt.Sprint(node)] {
			if !(contains(tovisit, v) || contains(visited, v)) {
				tovisit = append(tovisit, v)
			}
		}
	}

	fmt.Printf("Here the result:%v if we start from %d\n", visited, s)
	wg.Done()

}
