package main
import "fmt"

func main(){
	number(2);
	number("6");
}

func number(i int){
	switch i{
	case 1:
		fmt.Println("one");
	case 2:
		fmt.Println("two");
	case 3: 
		fmt.Println("three");
	case 4: 
		fmt.Println("four");
	default:
		fmt.Println("anything else");
	}
	
}