package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	for caseNum := 0; caseNum < k; caseNum++ {
		adjList := make(map[int][]int, 0)
		scanner.Scan()
		numNodes, _ := strconv.Atoi(scanner.Text())
		for i := 1; i <= numNodes; i++ {
			adjList[i] = make([]int, 0)
		}

		scanner.Scan()
		numEdges, _ := strconv.Atoi(scanner.Text())
		for i := 1; i <= numEdges; i++ {
			scanner.Scan()
			source, _ := strconv.Atoi(scanner.Text())
			scanner.Scan()
			target, _ := strconv.Atoi(scanner.Text())
			adjList[source] = append(adjList[source], target)
		}
		cycleNotFound := true
		for node := range adjList {
			if !isAcyclic(adjList, node) {
				cycleNotFound = false
				break
			}
		}

		if cycleNotFound {
			fmt.Println(1)
		} else {
			fmt.Println(-1)
		}
	}
}

func isAcyclic(adjList map[int][]int, start int) bool {
	stack := make([]int, 0)
	stack = append(stack, start)

	var s struct{}
	visited := make(map[int]struct{}, 0)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		found := false
		for _, adjNode := range adjList[node] {
			if adjNode == start {
				return false
			}
			if _, ok := visited[adjNode]; !ok {
				stack = append(stack, adjNode)
				visited[adjNode] = s
				found = true
				break
			}
		}
		// If no unvisited nodes are adjacent pop the stack
		if !found {
			stack = stack[:len(stack)-1]
		}
	}
	return true
}
