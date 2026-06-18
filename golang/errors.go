package main
import (
	"fmt"
	"errors"
)

func withdraw(balance int, amount int) (int,error){
	if amount > balance{
		return -1, insufficient
	} else if amount <= 0{
		return -1, invalid
	} else {
		return balance-amount, nil
	}
}

var insufficient = errors.New("insufficient funds")
var invalid = errors.New("invalid amount")

func main(){
	if sisa, err := withdraw(10, 0); err != nil{
		if errors.Is(err, insufficient){
			fmt.Println("insufficient funds")
		} else if errors.Is(err, invalid){
			fmt.Println("invalid amount")
		} else {
			fmt.Println("Your balance is %s", sisa)
		}
	}
	if sisa, err := withdraw(10, 15); err != nil {
		if errors.Is(err, insufficient){
			fmt.Println("insufficient funds")
		} else if errors.Is(err, invalid){
			fmt.Println("invalid amount")
		} else {
			fmt.Println("Your balance is %s", sisa)
		}
	}
	if sisa, err := withdraw(10, 5); err != nil {
		if errors.Is(err, insufficient){
			fmt.Println("insufficient funds")
		} else if errors.Is(err, invalid){
			fmt.Println("invalid amount")
		} else {
			fmt.Printf("Your balance is %d", sisa)
		}
	} else {
		fmt.Printf("Your balance is %d", sisa)
	}
}