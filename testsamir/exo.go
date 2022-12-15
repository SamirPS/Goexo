package main

import (
	"fmt"
	"sync"
)

var Graph = make(map[string][]int)

func main() {
	var wg sync.WaitGroup
	Graph["0"] = []int{1, 2}
	Graph["1"] = []int{2, 3}
	Graph["2"] = []int{4}
	Graph["3"] = []int{4}
	Graph["4"] = []int{2, 5}
	Graph["5"] = []int{3}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			BFS(Graph, n)
		}(i)

	}
	wg.Wait()

}

func BFS(Graph map[string][]int, s int) {
	tovisit := make(chan int, len(Graph))
	var visited = map[string]bool{}

	fmt.Printf("\nWe begin with the %v node \n", s)

	for i := range Graph {
		visited[i] = false
	}
	tovisit <- s
	visited[fmt.Sprint(s)] = true

	for len(tovisit) > 0 {
		node := <-tovisit
		fmt.Printf("%v ", node)
		for _, v := range Graph[fmt.Sprint(node)] {
			if !visited[fmt.Sprint(v)] {
				tovisit <- v
				visited[fmt.Sprint(v)] = true
			}

		}
	}
	close(tovisit)

}
