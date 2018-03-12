package main

import (
	"fmt"
)

func sqrt1(num int) int {
	var temp int
	var sqrtnum int
	temp = num / 2
	for i := 1; i < temp; {
		sqr := i * i
		if sqr < num {
			i++
		} else if sqr > num {
			sqrtnum = i - 1
		} else {
			sqrtnum = i
			break
		}
	}
	return sqrtnum
}
func main() {
	var m map[string]int 
	m=make(map[string]int)
	m["one"]=1
	m["two"]=2
	fmt.Println(m)
	var sum int = 0
	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Println(sqrt1(16))
}
