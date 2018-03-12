package main
import "fmt"

type node struct {
	num int
	next  *node
}
var first,last *node
func createnode(num int) {
	var  item node
	item.num=num
	insertnode(&item)
}
func insertnode(item *node ) {
	if first==nil {
		first=item
}else {
	for i:=first;i!=nil;i=i.next {
		if i.next==nil {
			i.next=item
			last=item
			break
}
}
}
}
func queue() {
for i:=first;i!=nil;i=i.next {
	fmt.Println(i.num)
}

fmt.Println("Queue finished")
}
func stack() {
	var a[12]int
	j:=0
	for i:=first;i!=nil;i=i.next {
		a[j]=i.num
		j++

}
	fmt.Println("Stack.....")
	for k:=j-1;k>=0;k-- {
	fmt.Println(a[k])
}
}
func main() {
	for i:=0;i<10;i++ {
	createnode(i)
}

queue()
stack()
}
