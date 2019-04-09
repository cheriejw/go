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
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	m, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		arr = append(arr, num)
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		if i != m-1 {
			fmt.Printf("%d ", binarySearch(arr, num))
		} else {
			fmt.Printf("%d\n", binarySearch(arr, num))
		}
	}
}

func binarySearch(arr []int, num int) int {
	result := bSearch(arr, num, 0, len(arr))
	if result != -1 {
		return result + 1
	}
	return -1
}

func bSearch(arr []int, num int, left int, right int) int {
	mid := left + ((right - left) / 2)
	if arr[mid] == num {
		return mid
	} else if left >= right {
		return -1
	} else if num > arr[mid] {
		return bSearch(arr, num, mid+1, right)
	} else { //n < arr[mid]
		return bSearch(arr, num, left, mid-1)
	}
}
