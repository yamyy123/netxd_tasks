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
		return
	}
    front.next=head
	head.prev=front
	head=front
}

func insert(num,pos int){
	var th *node =head
	if(pos<=0){
		addfront(num)
		return
	}
	if(pos>count){
		addback(num)
		return
	}
	count++

	for ;pos>1; {
        th=th.next
		pos--
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
		return
	}
	back.prev=tail
	tail.next=back
	tail=back
}

func deletefront(){
   count--
   if(head==nil){
	fmt.Println("the list is empty")
	return
   }
   if(head.next==nil){
	head=nil
	tail=nil
	return
   }
   head=head.next
   head.prev=nil
}

func deleteback(){
	count--
	if(head==nil){
		fmt.Println("the list is empty")
		return
	   }
	   if(head.next==nil){
		head=nil
		tail=nil
		return
	   }
	   tail=tail.prev
	   tail.next=nil
}

func deleterandom(pos int){
	if(pos==0){
		deletefirst()
		return
	   }
	   if(pos==count){
		deletelast()
		return
	   }
	count--
	if(head==nil){
		fmt.Println("the list is empty")
		return
	   }
	   if(head.next==nil){
		head=nil
		tail=nil
		return
	   }
	   var th *node=head
	   
	   for ;pos>0; {
             th=th.next
			 pos--
	   }
	   th.prev.next = th.next
	   th.next.prev=th.prev
}
 func fronttraversal(){
       var th *node=head
	   for ;th.next!=nil;{
          fmt.Println(th.val)
		  th=th.next
	   }
 }


 func backtraversal(){
	var th *node=tail
	for ;th.prev!=nil;{
	   fmt.Println(th.val)
	   th=th.prev
	}
 }

func main(){
 fmt.Println("choose from the options given below")
 fmt.Println("\n1.Insert in begining\n2.Insert at last\n3.Insert at any random location\n4.Delete from Beginning\n  
 5.Delete from last\n6.Delete the node after the given data\n7.Search\n8.Show\n9.Exit\n")
 var choice int
 fmt.Scanln(&choice)
 switch(choice){
 case 1:{
	var num int
	fmt.Println("enter the value to be added in the first")
	fmt.Scanln()
 }
 }
}