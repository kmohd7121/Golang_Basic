package main
import ("fmt")
type Vertex struct{
	Lat,Long float64
	
}
var m map[string]Vertex
func main(){
	m=make(map[string]Vertex)
	M["water park"] float64= Vertex{
		40.68433,
		-74.39967,
	}
	fmt.Println(m["water park"])
}