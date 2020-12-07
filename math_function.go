package main
import ("fmt"
"math")

func main(){
	mynum := 0.1
	if mynum == math.Pow(math.Sqrt(mynum),2){
		fmt.Println("these are the same")
	}else{
		fmt.Println("these are the different ")
	}
}