package main
import "fmt"

// a type of closure where a function returns a function
func item() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}

func main(){
	keren := item()
	fmt.Println(keren)
	fmt.Println(keren())
	fmt.Println(keren())
	fmt.Println(keren())
	fmt.Println(item)

	keren2 := keren // not creating a new instance but pointing that variable to the same address in memory, meaning if we print again later it will be +1 on that value i on that particular instance
	fmt.Println(keren2)
	fmt.Println(keren) // btw printing a function without calling it, we talking about the function not the result
	// thats why when printed, it will return the address in memory, function itu treated as nilali (first-class fuinction)
	// value of function is its own address in memory
	// first class function means function = value, can be treated as a normal data
	fmt.Println(keren2())
}