package main
import "fmt"

func zeroval(item int){
	item = 0
}

func zeropointer(item *int){ // item ini kan vriabel yang nyimpan alamat ke sebuah value bertipe integer
	*item = 0 // kita grab dulu valuenya berapaa, dengan pake *item di body fungsi, ngegrab value dari address tersebut lalu ktia jadiin 0
	// kalau langsung item = 0 kita ngerubah alamat bukan rubah valuenya bos
}

func main(){
	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)

	zeropointer(&i) // sebagai input ke function zeropointer adalah alamat dari variabel i atau addressnya
	// yang kita grab dengan &i yaitu kita ngegrab address dari  variabel i ini
	fmt.Println(i)
}