package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var fibMap map[int]int

func main() {
	fibMap = make(map[int]int)
	fibMap[0] = 0
	fibMap[1] = 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(fib(n))
	}
}

func fib(n int) int {
	if v, ok := fibMap[n]; ok {
		return v
	}
	fibMap[n] = fib(n-1) + fib(n-2)
	return fibMap[n]
}
