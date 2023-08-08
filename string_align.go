package main

import "fmt"

/*func swap (p1 ,p2 *byte){
     temp:=*p1;
    *p1=*p2;
    *p2=temp;
}*/

func main() {
    var s string
    fmt.Scanln(&s)
    var count int=0
    b:=[]byte(s)
    var lim int=len(b)
    for i:=0;i<lim;i++{
        count++
        if(count==3 && i!=lim-1){
            count=0
            //swap(&b[i],&b[i+1])
            b[i],b[i+1] = b[i+1],b[i]
            i++
        }
    }
    fmt.Println(string(b))
}