package main

import "fmt"

func main(){
	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			result:=i*j
			fmt.Printf("%d*%d=%d ",j,i,result)
		}
		fmt.Println()
	}
}
