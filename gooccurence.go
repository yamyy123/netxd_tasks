package main
import "fmt"

func main() {
   var str string
   fmt.Scanln(&str)
   var word string
   fmt.Scanln(&word)
   occurence :=0
   count :=0
   k:=0
   for i:=0;i<len(str);i++{
       if str[i]==word[k]{
           k++
           count++
           
       }else if str[i]==word[0]{
           count=0
           k=0
           count++
           k++
       }else{
           count=0
           k=0
   }
   if(count==len(word)){
       occurence++
       count=0
       k=0
   }
}
fmt.Println(occurence)
}