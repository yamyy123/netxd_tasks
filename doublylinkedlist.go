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

func insert(num,pos int){
	var th *node =head
	count++
	if(pos<=0){
		addfront(num)
		return;
	}
	if(pos>size){
		addback(num)
		return;
	}
	for ;pos>1; {
        th=th.next
	}
    temp:=&node{nil,num,nil}
	temp.prev=th
	temp.next=th.next
	th.next=temp
	temp.next.prev=temp//if we use like this we can able to change temp variables next address previous variable
	



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
   head.prev=nil
}

func deleteback(){
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
	   tail=tail.prev
	   tail.next=nil
}




func main(){

}