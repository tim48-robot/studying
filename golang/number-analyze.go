package main
import "fmt"

func main(){
	numberAnalyzer(5)
}


func numberAnalyzer(n int){
	for i:= 1; i<=n; i++{
		var even bool = (i % 2 == 0);
		fmt.Println("the number is ", i);
		if (even){
			fmt.Println("its an even number")
		} else {
			fmt.Println("its an odd number")
		}
	}

	if n:=9; n>0{
		fmt.Println("Angka cukup besar")
	} else if n<0 {
		fmt.Println("Angka terlalu kecil")
	}
}