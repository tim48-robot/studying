package main
import "fmt"



type person struct {
	item int
	name string
}
 // if we use a constructor function later then its going to be better, btw & (pointer) to strucrt when being printed later is a little bit different than other pointers, since it will be just like &{Rex 15} for example, and if we wanna print real address goota do something like "%p\n
 /* we also have automatic dereference in GO, 
 	sp := &s
	fmt.Println(sp.age)
	very possible even etho pointer.age should be illegal, 
	automatic dereference mostly working on field access in struct 

	an adress that we want to conver to value is with *, --> *person
	a value that we want the address of we use &





*/

func main(){
	dog := struct {
		age int 
		name string
	} {
		15, "Rex",
	} // anonymous struct, we dont need to write type animal struct or type dog struct whatever

	fmt.Printf("%p\n", &dog) // sometimes after % and before we declare what type we want to change it in in this case which is p or pointer, 
	// we can add # as more info, for example %#U, or %#x, but if we add it to this pointer its kinda pointless tho 
	fmt.Println(&dog)
}