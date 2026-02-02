package main

import "fmt"

func main() {
	s := make([]int, 0, 3)

	s = append(s, 1)
	s = append(s, 2)

	t := s
	s = append(s, 3)

	t[0] = 99

	fmt.Println("s:", s)
	fmt.Println("t:", t)
	fmt.Println("len(s):", len(s), "cap(s):", cap(s))
	fmt.Println(cap(t))
	t = append(t, 10)
	t = append(t, 15)
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(cap(t))
	fmt.Println(cap(s))
	
}
