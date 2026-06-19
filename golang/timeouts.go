// here we are going to implement timeouts in go uosing simple select by comparing time & an action that is supposedly could with external, but in this case im just using an example where we append a value after a time
package main
import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(2 * time.Second)
		c1 <- "Hello" // this wont be sent btw since its unbuffered, after 2 seconds goroutine wakes up
		// there is a sender that is from goroutine to the channel but there is no receiver?
		// that means its stuck there since unbuffered channel need both receiver & sender.
		// even tho we use goroutine, it wont be sent to the channel, since it will be sent if it detects receiver, it try going down the code but until very last code of main there is still no receiver since select is already terminated with timeout
	}()

	select {
		case cool := <- c1:
			fmt.Println("%s first", cool)
		case <- time.After(1*time.Second):
			fmt.Println("Timeout 1 Second") 
	}

	go func(){
		time.Sleep(2*time.Second)
		c2 <- "World"
	}()
	select {
		case cool := <- c2:
			fmt.Printf("%s first", cool)
			fmt.Println()
		case <- time.After(3*time.Second):
			fmt.Println("3 Second First")
	}
}

// and btw before the second go func to execute which is to wait for 2 seconds before assigning the value "World" to c2, 
// we need to wait for the select to happen to choose either one because select is synchronous