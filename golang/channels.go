package main
import "fmt"
// func main(){	
// 	messages := make(chan string)
// 	go func() {messages <- "hello"}() // the only goroutine here
//     msg := <- messages 
// 	fmt.Println(msg)
// }
// channel is how goroutine communicate? 
// so here, making a channel, then create a goroutine where hello string is being injected into the channel
// because go routing is concurrent and can run in the background, we go to the msg:= <- messages, since the code is in main and its synchoronnusly there we wait until the messages has a content, if yes, then the value is sent to msg and its printed


func main() {
    ch := make(chan string, 3)
    ch <- "a"
    ch <- "b"
    ch <- "c"
    fmt.Println(<- ch)
	
}