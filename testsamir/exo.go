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
	wg.Add(1)
	go BFS(Graph, 1)
	wg.Wait()

}

func contains(s *chan int, element int) bool {
	for v := range *s {
		if v == element {
			return true
		}
	}
	return false
}

func BFS(Graph map[string][]int, s int) {
	tovisit := make(chan int, len(Graph))
	visited := make(chan int, len(Graph))

	tovisit <- s
	for len(tovisit) > 0 {
		node := <-tovisit
		visited <- node
		for _, v := range Graph[fmt.Sprint(node)] {
			if !(contains(&tovisit, v) || contains(&visited, v)) {
				tovisit <- v
			}
		}
	}

	fmt.Printf("Here the result:%v if we start from %d\n", visited, s)
	wg.Done()

}
