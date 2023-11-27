package main

import "fmt"

func simCheck(matrix [][]int, n int) string {
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if matrix[i][j] != matrix[j][i] {
				return "no"
			}
		}
	}
	return "yes"
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([][]int, n)

	for i := range arr {
		arr[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	fmt.Println(simCheck(arr, n))
}
