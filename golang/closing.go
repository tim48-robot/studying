package main
import "fmt"


func main(){
	jobs := make(chan int, 1)
	done := make(chan bool)
	go func(){
		for {
			value, more := <- jobs
			if more {
				fmt.Printf("Received Job %d", value)
				fmt.Println()
			} else {
				fmt.Println("Receive All Jobs")
				done <- true
				return
			}
		}
	}()

	for i:=1; i<4; i++{
		jobs <- i
		fmt.Printf("Sent Job %d", i)
		fmt.Println()
	}
	close(jobs)
	fmt.Println("Sent All Jobs")
	<-done


	_, dones := <- jobs
	fmt.Printf("should be false -> %t", dones)
	fmt.Println()
}

// seeems easy but the sent and receive is so weird at the output if you see, its because go scheduler can context switch anytime and it mustn't go from for loop togo function then for loop then go function
// go function runs in the background so for loop can run twice then go routine or maybe for loop thrice then go routine thrice, 
// more also could be true when the channel is closed if there is still value, more is only false if the chhanel is closed + there is no value.
