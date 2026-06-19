// channel directions, here it is even simpler
// how i can see it is that only 2 operation are allowed.
/*
1. send to channel.
2. receive from channel.

don't misunderstood receive from channel as send from channel. receive & send need to be understood properly
here i have a function pings and pongs where pings is sending message to a channel and pongs is sending message to another channel.

*/


// chan <-, send to channel (first operation)
package main
import (
	"fmt"
)
func ping(pings chan<- string, message string){
	pings <- message
}

func pong(pings <-chan string, pongs chan<- string){
	msg := <-pings
	pongs <- msg
}

func main(){
	pings := make(chan string) // created as bidirectional, but can be converted to directional when passed to a certain function like ping or pong, but the thing is once inside the function in the parameter you already said what direction it is so inside the function you cant turn the directional backwards again. this violates all principle of type safety of golang.
	pongs := make(chan string) // if we wanna create as directional we need to use something liek pongs := make (chan<- string). useless tho, we better think about the rules when inside the function so that we can change it according to our needs.
	go ping(pings, "hello") // if we create it in main as directional, the chanenl use is literally only limited to that specific function for example receive or send only which is basically useless. type safety in function is enough
	go pong(pings, pongs)
	fmt.Println(<-pongs)
}