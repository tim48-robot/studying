// a simple non blocking channel operations, it is to minimize the chance of deadlock, for example
// we try to send or receive to a non buffer channel with a non goroutine way, then deadlock is for sure.
// we can get rid of the chance of deadlock basically by saying, if you go non goroutine + non buffer.
// you can still have no deadlock basically by saying no value? then we go with default, which is good for select

package main
import "fmt"
func main(){
	msg := make(chan string)
	signal := make(chan bool)

	select {
	case message := <- msg:
		fmt.Printf("this %s is very important", message)
    default:
		fmt.Println("no message yet")
	}

	select {
	case message := <- msg:
		fmt.Printf("this %s is very important", message)
	case <- signal:
		fmt.Printf("there is a signal")
    default:
		fmt.Println("no message nor signal yet")
	}

}

// perfect summary:
/*
unbuffered, gotta have both sender & receiver in the same time, impossible if synchronous, you need go routine. IF synchoronous, probably will go to default
buffered, you dont need both sender & receiver in the same time, just put it in the buffer first & receiver/sender can take or sent it later
*/