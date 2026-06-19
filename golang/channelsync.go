package main
import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Println("Working....")
	time.Sleep(time.Second)
	fmt.Println("done!")
	
	done <- true
}

func main(){
	done := make(chan bool, 1)
	go worker(done)
	<- done // main wont terminate except the var done has a value, and then we get rid of that value. only after do we terminate main.
}