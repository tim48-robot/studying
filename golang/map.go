package main

import (
	"fmt"
)

func main(){
	m := map[string]int{
		"a":1,
		"b":2,
	}
	fmt.Println(saferead(m, "b"))
	fmt.Println(saferead(m, "c"))

	m2 := map[string]int{
		"apple":5,
		"banana":3,
		"orange":5,
	}
	fmt.Println(reverseindex(m2))

	scores := map[string]int{
    	"alice": 0,
	}
	zerovalue(scores, "alice")
	zerovalue(scores, "kazuya")
	zerovalue(scores, "kazama")

	m3 := map[string]int{
    "a": 1,
    "b": 0,
    "c": 3,
	}

	deleteoniterate(m3);
}


func saferead(m map[string]int, key string)int{
	_, exist := m[key]
	if exist == true{
		return m[key]
	} else {
		return -1
	}
}

func reverseindex(m map[string]int) map[int][]string{
	result := make(map[int][]string)
	for key, value := range m {
		result[value] = append(result[value], key)
 	}
	return result;
}

func mergemaps(a, b map[string]int) map[string]int{
	result := make(map[string]int)
	for key,value := range a{
		result[key] = value;
	}
	for keyb,valueb := range b{
		result[keyb] = valueb;
	}
	return result;
}

func zerovalue(m map[string]int, name string){
	if _, exist := m[name]; exist{
		fmt.Println(name, "has a score")
	} else {
		fmt.Println(name, "doesnt has a score")
	}
}

func deleteoniterate(m map[string]int){
	for key,value := m {
		if m[key] == 0{
			delete(m, key)
		}
	}
}