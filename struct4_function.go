package main

import "fmt"
 type animal struct{
	 name string 
	 origin string
 }
 type bird struct{
	 animal
	 speedKH int
	 canfly bool
 }
 func main(){
	 b:=bird{}
	 b.name="emu"
	 b.origin="Australia"
	 b.speedKH  =48
	 b.canfly = false
	 fmt.Println(b)
 }
