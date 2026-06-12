/*
 in go, we have ~[]E for example, this is to let not only the hardcoded string type.
 but lets say we have a struct type MyString []string, that supposedly underlying is []string 
 but its a custom type, with ~ it is allowed.

 as for comparable, meaning that we only accept those datatypes that is comparable.
 seems that for generics we can use both [] and () at the same time beside the function.
 what i see is that [] is more for defining rules and () is for telling in runtime, variable name example & must follow variable type

*/

/*
IMPORTANT!
difference between * and & in its basics.
* will be used only at type, or could be for derefrencing an address. lets say p := &int,
we can get the value by doing *&(int), which is taking the value.
or when stating that a variable should have a pointer type. (not raw value but pointer)  

*/

// now imma try implement a stack btw
// last in first out, we can make first out basically cutting the tail, singly linked list should suffice
// need to have the function push pop and contains btw
package main
import "fmt"

type Stack[T comparable] struct{
	head, tail *node[T]
}

type node[T comparable] struct{
	next *node[T] // btw this just means that node[T] is like a type that 
	val T
}

func (lnkdlst *Stack[T])Push(v T){
	if lnkdlst.head == nil || lnkdlst.tail == nil{
		lnkdlst.head = &node[T]{val: v}
		lnkdlst.tail = lnkdlst.head
	} else {
		newNode := &node[T]{val: v, next:lnkdlst.head} // next not told, so next would be nil.
		lnkdlst.head = newNode
	}
	// lnkdlst is a pointer, auto dereference btw so we can use the value without needing to do &
	// btw nil can be done on non comparable elements right?
}

func (lnkdlst *Stack[T]) Pop() (T, bool){
	if lnkdlst.head != nil{
		tempval := lnkdlst.head.val
		lnkdlst.head = lnkdlst.head.next
		return tempval, true
	} else {
		var zero T
		return zero, false
	}
} 

func (lnkdlst *Stack[T]) Contains(v T) bool{
	for i := lnkdlst.head; i != nil; i = i.next {
		if i.val == v{
			return true
		}
	}
	return false
}

func main(){
	newItem := Stack[string]{}
	newItem.Push("a")
	newItem.Push("b")
	newItem.Push("c")
	newItem.Push("d")
	fmt.Println(newItem.Pop())
	fmt.Println(newItem.Pop())
	fmt.Println(newItem.Contains("b"))
	fmt.Println(newItem.Contains("a"))


	newItem2 := Stack[int]{}

	newItem2.Push(1)
	newItem2.Push(2)
	newItem2.Push(3)
	newItem2.Push(4)

	fmt.Println(newItem.Pop())
	fmt.Println(newItem.Pop())
	fmt.Println(newItem.Pop())
	fmt.Println(newItem2.Pop())
	fmt.Println(newItem2.Pop())
	fmt.Println(newItem2.Pop())
	fmt.Println(newItem2.Pop())
	fmt.Println(newItem2.Pop())
	fmt.Println(newItem2.Pop())


}
