package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type node struct {
	rollno string
	name   string
	token  string
	next   *node
}

var first *node

func createnode(data string) {
	var item node
	parsedString := strings.Split(data, ",")
	item.rollno = string(parsedString[0])
	item.name = string(parsedString[1])
	item.token = string(parsedString[2])
	item.next = nil
	insertnode(&item)
}
func insertnode(item *node) {
	if first == nil {
		first = item
	} else {
		for i := first; i != nil; i = i.next {
			if i.next == nil {
				i.next = item
				break
			}
		}
	}
}

func printdata() {
	for i := first; i != nil; i = i.next {
		fmt.Printf("ROllno:%s,name:%s,token:%s \n", i.rollno, i.name, i.token)
	}
}
func main() {
	content, err := ioutil.ReadFile("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	parsed := strings.Fields(string(content))
	for i := 0; i < len(parsed); i++ {
		createnode(parsed[i])
	}

	printdata()
}
