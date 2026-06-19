package main
import (
	"fmt"
)

func f (from string){
	for i:= range 100{
		fmt.Println(from, ":", i)
	}
}

func main(){
	f("hello")
	go f("routine1")
	go f("routine2")
}

// go routine concurrent & async
// concorurent means many tasks can be done at once, 
// normally concurent + async or concurent + paralel, paralel is a better version than async, mostly need multi-core CPU, as for async -> you can  do more than one thing at once, meaning that you can switch to another process while waiting.
