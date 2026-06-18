package main
import (
	"fmt"
	"iter"
)



func main(){	
	for v := range Numbers(5){
		fmt.Println(v)
	}
	
	for v := range Filter(Numbers(10), func(n int) bool{
		return n%2 == 0
	}) {
		fmt.Println(v)
	}

}

func Numbers(n int) iter.Seq[int]{
	return func(yield func(int) bool){
		for i:=1; i<=n; i++{
			if !yield(i){
				return
			}
		}
	}
}


// somehow yield(item) gives the fvalue to the outer part of the program, in this case even if checking !yield(item) when trying to return false & terminate
// this behaviour is similar to when you try to pass a function to a variable.
// somewhat like this, var cool = breakable()
// you would pass the breakable function to the variable 
func Filter(r iter.Seq[int], condition func(int) bool) iter.Seq[int]{
	return func(yield func(int) bool){
		r(func (item int)bool{
			if condition(item){
				if !yield(item){
					return false
				}
			}
			return true
		})
	}
}
