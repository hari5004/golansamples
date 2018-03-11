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
	for i:=first ;i!=nil;i=i.next {
		fmt.Println(i.num)
}
}
func main() {
for i:=0;i<10;i++ {
	createnode(i)
}
printnode()
}


