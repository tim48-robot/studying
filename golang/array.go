package main
import "fmt"

func main(){
	var nums [4]int
	nums[0] = 10
	nums[3] = 40
	for i:=0; i<len(nums); i++{
		fmt.Print(nums[i], " ")
	}
	fmt.Println("")
	fmt.Println(nums[3])
	fmt.Println(len(nums))
	d2()
}

func d2(){
	var d2var [2][2]int

	var counter int = 1;
	for i := range 2{
		for j := range 2{
			d2var[i][j] = counter;
			counter++;
		}
	}
	fmt.Println(d2var)
}
