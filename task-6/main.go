package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Replace(str, " ", "", -1)

	mapa := make(map[byte]int)

	for i := range str {
		mapa[str[i]] = 1
	}

	fmt.Println((len(mapa)) - 2)
}
