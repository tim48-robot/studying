// select is a way for us to wait for more than 1 channels to have a value before the main program is terminated
// 

package main
import (
	"fmt"
	"time"
)

func main(){
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func(){
		time.Sleep(1 * time.Second)
		channel1 <- "Hello"
	}()
	go func(){
		time.Sleep(2 * time.Second)
		channel2 <- "World"
	}()

	for range 2 {
		select {
		case msg1 := <- channel1:
			fmt.Println(msg1)
		case msg2 := <- channel2:
			fmt.Println(msg2)
		}
	}
}
// but well for range n, the n value here corresponds to the amonut of value that is sent from the channel, if we have 2 channel & 1 is buffered with many values, it is possible that channel1 send values multiple times to msg1 & msg2 doesn't print anything, especially if the time that they wait for is different.
// altho in this case both is unbuffered which is good!