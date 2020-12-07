package main
import "fmt"
func main(){
	adoctor:=struct{name string}{name:"jone petter "}
	anotherDoctor := adoctor
	anotherDoctor.name = "tom breaker"
	fmt.Println(adoctor)
	fmt.Println(anotherDoctor)
}