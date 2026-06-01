package main
import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fibonacci(n-1)
	}
}


func main(){
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(12))
	fmt.Println(fibonacci(7))


	var fib func(n int) int 

	fib = func(n int) int { // fib is a variable where value is a function that has an input of integer and returns an integer
		/* its very simplistic here, but the key part is that you cant do
		fib := func(n int) int etc etc at once, since golang check right value first before assigning to fib
		but if right value calls fib where it doesnt exist yet, then its not possible
		*/

		if (n == 0) {
			return 1 
		} else {
			return n * fib(n-1)
		}
	} // factorial btw not fibonacci

	fmt.Println(fib(4))


// somehow for a variable that lets say it is outside the anonymosu function, and then in side the function it uis being called & appended with a value, i would say that the specific varaible is in a global state that can be modified. since it is global and modified from local.


	// ANOTHER CASE
	funcs := []func(){}
	for i := 0; i < 3; i++ {
		i := i  // creating a new i copy every iteration, proven by address
		fmt.Println(&i) // later if called per item, specific funcs[0] to funcs[2] will print 0,1,2
		// since it doesnt use global i which is == 3 after last for loop is executed 
		funcs = append(funcs, func() {
			fmt.Println(i)
		}) // btw here three different instance function with different i, since the code is identic then the address funfction is same, but the value address is different
	}
	funcs[0]()
	funcs[1]()
	(funcs[2]())

}