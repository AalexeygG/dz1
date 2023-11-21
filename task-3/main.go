package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	slice := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&slice[i])
	}
	for i, j := 0, N-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	for i, j := 1, N-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println(slice)
}
