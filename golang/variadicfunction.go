// here we are working with variadic function in golang
// variadic function is you get the input as like multiple values or input of a certain data typepe
// if you input it as an array, gotta have to brea kti first using ... for example
// treated quite like an array inside? because when you input it as a like some elemnt of a type
// then inside its treated like an array, well we can check using "reflect" packag

package main

import (
	"fmt"
	"reflect"
)

func keren(nums ...int){
	fmt.Print(nums, " ")
	total := 0
	length := 0

	fmt.Println(reflect.TypeOf(nums))

	for _, num := range nums {
		total += num
		length ++
	}

	fmt.Println(total)
	fmt.Println(length)
}

func main(){
	keren(1,2,4,16)
	keren(1,6,1,8)
	keren(1)
	keren([]int{1,3,1,7,1,18,18}...)
}

// just so i remember, you can run right away using go run {filename}.go
// or build first like go build {filename}.go after that executing using either {filename}.exe or ./{filename} for linux & mac