package main

import "fmt"

type queue struct {
	front,rear,size int
	capacity int 
	arr[]int
}

func Newqueue(capacity int)(*queue){
	return &queue{
		front :0,
		rear:-1,
		size:0,
		capacity:capacity,
		arr:make([]int,capacity),

	}
}

func(q *queue) enqueue(item int)(bool){
	if(q.size==q.capacity){
		fmt.Println("queue is full")
		return false
	}
	q.rear=((q.rear+1)%q.capacity)
	q.arr[q.rear]=item
	q.size++
	return true
}

func(q *queue) dequeue()(int,bool){
	if(q.size==0){
		fmt.Println("queue is empty")
		return 0,false
	}
	item :=q.arr[q.front]
	q.front=((q.front+1)%q.capacity)
	q.size--
	return item,true
}

func main(){
	q :=Newqueue(5)
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.enqueue(40)
	q.enqueue(50)
	//q.enqueue(60)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	//fmt.Println(q.dequeue())

}