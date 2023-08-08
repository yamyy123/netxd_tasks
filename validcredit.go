package main

import "fmt"

func evenfunc(temp int32) (int32){
	
	if temp%10==temp{
         return temp
	}else {
          var e int32=0;
		  for i:=0;temp!=0;i++{
			e=e+(temp%10)
			temp=temp/10
		  } 
		  return e
	}
}

func checkvalid(str string) (int32){
	k:=1
   b:=[]int32(str)
     if len(str)>=13 && len(str)<=16{
		k=1
	 }else{
		return 1
	 }
	 for i:=0;i<len(str);i++{
		if b[i]>=48 && b[i]<=57{
             k=1
			 b[i]=b[i]-48
		}else{
			return 1
		}
	 }
	 fmt.Println()//
	 if  (b[0]==4)||(b[0]==5)||(b[0]==6)||(b[0]==3&& b[1]==7){
		k=1
	 }else{
		return  1
	 }
	 var oddsum int32=0
	 var evensum int32=0
	 var temp int32=0
	 count:=1
	 for i:=len(b)-1;i>=0;i--{
		if count%2!=0{
			oddsum = oddsum+(b[i])
		}else{
            temp=b[i]
			temp=temp*2
			evensum=evensum+evenfunc(temp)	
		}
		count++
	 }
	 fmt.Println(evensum)
	 fmt.Println(oddsum)
	 valid:=evensum+oddsum
	// fmt.Println(valid)
	 if valid%10==0 && k==1{
		return 0
	 }else{
		return 1
	 }
}

func main(){
	var str string
	fmt.Scanln(&str)
	if checkvalid(str)==0 {
		fmt.Println("The card is valid")
	}else{
		fmt.Println("the card is invalid")
	}
}