package main
import "fmt"

func main(){
	messages := make(chan string, 2)

	go func() {
		messages <- "hello"
		messages <- "world"
		messages <- "i love my right, can we know each other"
	} ()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// buffer lets you store in the channel n amount of message of type x
// here we dont use go routien since its still 2, if we declare n amount and put the message >n amount we need to use goroutine.
// so that it wont be deadlock, but you gotta know that the channel size is still n and if you want to free the size you gotta take the message out of the channel & then the next message will be inserted inside