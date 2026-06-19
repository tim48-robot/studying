/*
data type of struct is that we define
data type of interface is that we would normally use the function that is the preconditon of being that interface
in this case, 
type error interface {
	Error() string
}

anyways im going to make a custom error using that stuff?
*/
package main
import (
	"fmt"
	"errors"
)

type argError struct{
	data int
	argument string
}

// here i dont understand why we need pointer while copy is safer, altho maybe overhead copy will not take much memory
// but i guess go convention is just like that 
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s",e.data, e.argument)
}

func divide(a int, b int)(int, error){
	if b == 0 {
		return -1, &argError{b, "cannot divide by zero"}
	} else {
		return a/b, nil
	}
}

func main(){
	_, err := divide(1,5) // err type is error btw, since return of divide is error
	// thats why gotta open the interface using the errors.AsType
	var ae *argError
	if errors.As(err, &ae){
		fmt.Println(ae.data, ae.argument)
	} else {
		fmt.Printf("This is not an argError")
		fmt.Println()
	}

	_, err = divide(1,0)
	if errors.As(err, &ae){
		fmt.Println(ae.data, ae.argument)
	} else {
		fmt.Printf("This is not an argError")
	}
}





/// VERY INTERESTING