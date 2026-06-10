package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct{
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64{
	return r.width * r.height
}

func (r rect) perim() float64{
	return 2 * (r.width + r.height)
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius 
}

func (c circle) perim() float64{
	return 2 * math.Pi * c.radius
}

func interfaceGeo(g geometry){
	fmt.Println(g.area())
	fmt.Println(g.perim())

	// here you do type assertion using interface by the text g.(circle) only useful on an interface
	if _, ok := g.(circle); ok {
		fmt.Println("this one is a circle")
	} else {
		fmt.Println("this one is a square")
	}
}

// func main(){
// 	c1 := circle{8}
// 	r1 := rect{10, 5}
// 	interfaceGeo(c1)
// 	interfaceGeo(r1)
// }































type Notifier interface {
	send(string) 
}

type EmailNotifier struct {
	message string
}

type SMSNotifier struct {
	message string
}

type UserService struct {
	notifier Notifier
}


func (en EmailNotifier) send(message string) {
	fmt.Println("Email sent: " + message)
}

func (sn SMSNotifier) send(message string){
	fmt.Println("SMS sent: " + message)
}


func NewUserService(n Notifier) UserService {
	return UserService{notifier: n}
}

func (us UserService) Register(username string) {
	us.notifier.send("Welcome, " + username)
}

func main(){
	smsnot := SMSNotifier{}
	emailnot := EmailNotifier{}
	NewUserService(smsnot).Register("Budi")
	NewUserService(emailnot).Register("Budi")
}
