package main

import (
	"bytes"
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
	/*
		tovisit is the channel of node i need
		to check.
		Visited is a dict which permit me to
		know, if i have already visited the node
	*/

	tovisit := make(chan int, len(Graph))
	var visited sync.Map
	b := &bytes.Buffer{}
	var wg sync.WaitGroup
	defer fmt.Println(b)

	fmt.Fprintf(b, "We begin with the node %v: \n", s)

	/*
		init the visited dict, false for all node
		except the departure node
	*/
	for i := range Graph {
		visited.Store(i, i == fmt.Sprint(s))

	}

	tovisit <- s
	for n := range tovisit {
		if n == -1 {
			wg.Wait()
			close(tovisit)
			break
		}
		fmt.Fprintf(b, " %d ", n)
		wg.Add(1)
		go func(Graph map[string][]int, visited sync.Map, tovisit chan int, node int) {
			defer wg.Done()
			searchNeighbours(Graph, visited, tovisit, node)
		}(Graph, visited, tovisit, n)
	}

}

func searchNeighbours(Graph map[string][]int, visited sync.Map, tovisit chan int, node int) {
	c := -1

	for _, v := range Graph[fmt.Sprint(node)] {
		result, _ := visited.Load(fmt.Sprint(v))
		if !result.(bool) {
			c += 1
			tovisit <- v
			visited.Store(fmt.Sprint(v), true)
		}

	}
	if c == -1 {
		tovisit <- c
	}

}
