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
	var visited = map[string]bool{}
	b := &bytes.Buffer{}
	defer fmt.Println(b)
	defer close(tovisit)

	fmt.Fprintf(b, "We begin with the node %v: \n", s)

	/*
		init the visited dict, false for all node
		except the departure node
	*/
	for i := range Graph {
		if i == fmt.Sprint(s) {
			visited[i] = true
		}
		visited[i] = false
	}

	tovisit <- s

	for len(tovisit) > 0 {
		node := <-tovisit
		fmt.Fprintf(b, "%v ", node)
		/* for each successor of node:
		- if we not have visited this successor :
			- we add it to the channel
			- and change the value of the visited map to true
		*/
		for _, v := range Graph[fmt.Sprint(node)] {
			if !visited[fmt.Sprint(v)] {
				tovisit <- v
				visited[fmt.Sprint(v)] = true
			}

		}
	}

}
