package main
import "fmt"

type rect struct{
	width, height, luas int
}

// anyways this mean that this can only be called by an object that is rect typed, cannot be done by something with int typed or anything else
func (r rect) keliling() int{
	return 2*r.width + 2*r.height
}

func (r *rect) calcArea() {
	r.luas = r.width * r.height
}

func (r rect) calcAreaNoPointer() { // this would behave weird since you sent a copy or duplicate of an object with the same value, and 
	r.luas = r.width * r.height
}

func main(){
	r := rect{width:10, height:5}
	fmt.Println(r.keliling())
	r.calcArea()
	// underlying nya karena calcArea itu butuh address dari asalnya, butuh pointer. 
	// jadi underlying this will be (&r.calcArea()), 
	// anyway, the key decision is from the function receiver, does it receive a pointer or an instsance of a struct in this case?
	// since it receives a pointer, then it will get the address of the instance r and then change that sepcfiic instance in that address.
	// 
	fmt.Println(r.luas)

	r2 := rect{width:10, height:5}
	r2.calcAreaNoPointer()
	fmt.Println(r2.luas) // 
}