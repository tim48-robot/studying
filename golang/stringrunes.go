package main

import "fmt"

func main() {
	s := "สวัสดี"
	fmt.Println(s[0]) // nilai dari byte yang by defaultnya decimal, keocuali kita bilang mau dihex atau dimana 
    // const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    // fmt.Print("Println: ")
    // fmt.Println(sample)

    // fmt.Println("Byte loop:")
    // for i := 0; i < len(sample); i++ {
    //     fmt.Printf("%x ", sample[i])
    // }
    // fmt.Printf("\n")

    // fmt.Println("Printf with %x:")
    // fmt.Printf("%x\n", sample)

    // fmt.Println("Printf with % x:")
    // fmt.Printf("% x\n", sample)

    // fmt.Println("Printf with %q:")
    // fmt.Printf("%q\n", sample)

    // fmt.Println("Printf with %+q:")
    // fmt.Printf("%+q\n", sample)


	// const nihongo = "日本語\xbd\xae"
    // for index, runeValue := range nihongo {
    //     fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    // }
}

/* string is normally trated as a []byte except for some cases where it is treated as a rune if it is converted to a rune or if it is iterated on a range 
lets  say that when iterating trhough a string using in range, there is cases where a certain rune is equal to 1 byte or more than 1 byte
if it is equal to 1 byte, 1 rune = 1 byte, and therefore the output would be like byte output
but if 1 rune has > 1byte, then the output is still on decimal but the decimal equivalent of that specific rune
lets say that the thai character ส which is in UTF-8 but not ASCII, so it needs > 1 byte,
when printed using for in range 
for idx, r := range s {
    fmt.Println(idx, r) // → 0 3626 (code point full of ส is U+0E2A).
    break
}
index will be 0 well yeah because rune is like that excepte byte where index may be 0, then 3 then 6 etyc
in this case rune is 
'ส' is a rune literal that is equivalent of the value 3626

Byte e0:

Hex: e0
Desimal: 224

Rune ส:

Unicode: U+0E2A
Hex: 0e2a
Desimal: 3626

*/

