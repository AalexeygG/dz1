package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	for i := range str { //только если используются латинские символы
		if str[i] == '1' {
			fmt.Print("one")
		} else {
			fmt.Print(string(str[i]))
		}
	}
}
