package main

import "fmt"

func main(){
	aDoctor:= struct{name string 
		age int
		comapny string}{name:"Jone Pertwer", age:25 ,comapny:"kloudone"}
	fmt.Println(aDoctor)
}