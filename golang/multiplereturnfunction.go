package main
import "fmt"


func plus() (int, int){
	return 3,7
}



func main(){
	a, b := plus()
	_, c := plus()

	// why must we use _? since in golang something that is declared but not unused will result in runtime error
	// thats why we need to throw the first value, and we cant just append it to a variable, but append it to _ since it has no value & used only for throwing unused subset of values 
	fmt.Println(c)
	fmt.Println(plus())	
}