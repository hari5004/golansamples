package main
import "fmt"
type node struct {
	num int
	next *node
}
var first *node
func createnode(num int){
	var p node
	p.num=num
	p.next=nil
//	fmt.Println(p.num)
	insertnode(&p)
}
func insertnode(p *node){
	if first==nil {
		first=p
}else {
	for i:=first;i!=nil;i=i.next {
		if i.next==nil {
			i.next=p
			break
}
}
}
}
func printnode() {
	fmt.Println(first.num)
	fmt.Println(first.next.num)
}
func main() {
createnode(1)
//fmt.Println(first.num)
createnode(2)
printnode()
}


