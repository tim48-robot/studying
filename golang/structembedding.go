/*
here i am going to study more about struct embedding.
embedding --> putting another struct as a field without a name into a new struct.
field only , no var name
*/

package main
import "fmt"


type innerStruct struct {
	num int
	str string
}

type cooliguess struct{
	innerStruct 
	number string
}

func (name innerStruct) cool() string{
	return name.str
}

func main(){
	var innerstruct = innerStruct{5, "kerenSTIN"}
	var item = cooliguess{innerstruct, "10"}
	fmt.Println(innerstruct.cool())
	fmt.Println(innerstruct.num)
	fmt.Println(innerstruct.str)
	fmt.Println(item.cool())
	fmt.Println(item.str)
	fmt.Println(item.num)

	


	// main function of struct embedding is only here really, calling a function without having to specficy item.innerStruct.cool(). Well i guess its kinda cool that it can do this
	//basically the difference with usual namned field like BOS innerstruct in cooliguess is that you have both method promotion Y& fields promotion
	// in the very outer part of the struct in this case is cooliguess, you can call cooliguess.cool()
	// or item.cool() or item.num and item.str right away. by not using a named field.	
	// by the weay must be done on a nmaed type meaning a type that we made from a certain type like int float64 or float 32 etc, but we made the type and named it ourselves.
	// else you cant use this no variabel name only type on otehr except named type
}