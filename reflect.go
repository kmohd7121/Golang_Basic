package main
import ("fmt" 
  "reflect")

type animal struct{
	name string
	required max:"100",
	origin string

}
func main(){
	f:=reflect.TypeOf(animal{})
	field,_:=f.FieldByName("name")
	fmt.Println(field.Tag)

}
