package main

import "fmt"
type node struct{
	 val int
	 prev *node
	 next *node

}

var head *node=nil;//these three are global variables
var tail *node=nil;
var count =0

func addfront(num int){
	count++
	front:=&node{ nil, num, nil}
	if head==nil{
		head=front
		tail=front
		return;
	}
    front.next=head
	head.prev=front
	head=front
}

func addback(num int){
	count++
	back:=&node{nil,num,nil}
	if head==nil{
		head=back
		tail=back
		return;
	}
	back.prev=tail
	tail.next=back
	tail=back
}

func deletefront(){
   count--
   if(head==nil){
	fmt.Println("the list is empty")
	return;
   }
   if(head.next==nil){
	head=nil
	tail=nil
	return;
   }
   head=head.next

}
func main(){

}