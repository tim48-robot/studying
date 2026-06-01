package main
import "fmt"


func plus (a int) int{
	return a + a
}

func plusPlus (a, b int) int {
	return a + b
}

func main(){
	fmt.Println(plus(5))
	fmt.Println(plusPlus(4,5))


}