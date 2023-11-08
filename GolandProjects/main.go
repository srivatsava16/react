package main
////
////import "fmt"
////
////func main() {
////	for i := 0; i < 6; i++ {
////		for j := 0; j < 5; j++ {
////			if i == 0 && j == 2 {
////				fmt.Println("\t\t*")
////			}
////			else if i==1&&j==2 {
////
////			}
////		}
////	}
////}
////package main
////
////import (
////	"fmt"
////)
////
////func main() {
////
////	ch := make(chan int, 2)
////	ch <- 2
////	ch <- 1
////	fmt.Println(<-ch + <-ch)
////}
//
////
////
////package main
////
////import (
////	"fmt"
////	"time"
////)
////
////func A(a int) {
////	for i := 0; i < a; i++ {
////
////		fmt.Println(i)
////	}
////}
////
////func B(c int) {
////	for i := 0; i < c; i++ {
////
////		fmt.Println(i)
////	}
////}
////func main() {
////	go A(1)
////	time.Sleep(2 * time.Second)
////	B(11)
////}
//
////package main
////
////import (
////	"fmt"
////	"time"
////)
////
////func main() {
////	ch := make(chan int)
////
////	ch <- 23
////	a := func(ch chan int) {
////
////		fmt.Println(234 + <-ch)
////	}
////	time.Sleep(2 * time.Second)
////
////	go a(ch)
////}
//
////
////package main
////
////import (
////	"fmt"
////)
////
////func a(ch chan int) {
////	fmt.Println(<-ch)
////}
////func main() {
////	ch := make(chan int)
////
////	go a(ch)
////	//time.Sleep(1 * time.Second)
////
////	ch <- 11
////
////}
//
////package main
////
////import "fmt"
////
////func myfunc(ch1 chan int, ch2 chan int) {
////
////	fmt.Println(<-ch1 + <-ch2)
////}
////func main() {
////
////	ch1 := make(chan int)
////	ch2 := make(chan int)
////	go func() {
////		ch1 <- 23
////		ch2 <- 26
////		a,b:=ch1,ch2
////
////	}()
////	go myfunc(a,b)
////
////	//myfunc(ch1, ch2)******************************
////
////}
//
////package main
////
////import "fmt"
////
////func main() {
////	ch := make(chan int, 2)
////	ch <- 12
////	ch <- 11
////	fmt.Println(<-ch + <-ch)
////
////}
////package main
////
////import "fmt"
////
////func main() {
////	ch := make(chan int, 2)
////
////	var a, b int
////	fmt.Scanln(&a, &b)
////	ch <- a
////	ch <- b
////	fmt.Println(<-ch + <-ch)
////
////}
////
//package main
//
//import (
//	"fmt"
//)
//
//func myfunc(ch1 chan int, ch2 chan int) {
//
//	fmt.Println(<-ch1 + <-ch2)
//	fmt.Println(<-ch1 - <-ch2)
//
//}
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go func() {
//		var a, b int
//		fmt.Println(&a, &b)
//		ch1 <- a
//		ch2 <- b
//	}()
//	go myfunc(ch1, ch2)
//}
//
////
////package main
////
////import (
////"fmt"
////"time"
////)
////
////func a() {
////	ch := make(chan int)
////
////	go func() {
////		ch <- 11
////	}()
////	fmt.Println(<-ch)
////
////}
////func main() {
////	go a()
////	time.Sleep(3*time.Second)
////
////}
////package main
////
////import "fmt"
////
////func sum(ch1 chan int, ch2 chan int) {
////	fmt.Println(<-ch1 + <-ch2)
////}
////func sub(ch1 chan int, ch2 chan int) {
////	fmt.Println(<-ch1 - <-ch2)
////}
////func ss() {
////	ch1 := make(chan int)
////	ch2 := make(chan int)
////
////	var a, b int
////	fmt.Println("Enter 2 numbers")
////	fmt.Scan(&a, &b)
////	ch1 <- a
////	ch2 <- b
////}
////func main() {
////	ch1 := make(chan int)
////	ch2 := make(chan int)
////
////	go ss()
////
////	var operator string
////	fmt.Println("enter operator")
////	fmt.Scanln(&operator)
////	if operator == "+" {
////		sum(ch1, ch2)
////	} else {
////		sub(ch1, ch2)
////	}
////}
////
////package main
////
////import (
////	"fmt"
////	// "time"
////)
////
////func sum(ch1 chan int, ch2 chan int) {
////	fmt.Println(<-ch1 + <-ch2)
////}
////func sub(ch1 chan int, ch2 chan int) {
////	fmt.Println(<-ch1 - <-ch2)
////}
////func mul(ch1 chan int, ch2 chan int) {
////	fmt.Println(<-ch1 * <-ch2)
////}
////func div(ch1 chan int, ch2 chan int) {
////	a, b := <-ch1, <-ch2
////	fmt.Println(a / b)
////
////}
////func main() {
////	ch1 := make(chan int)
////	ch2 := make(chan int)
////
////	go func() {
////		fmt.Println("Enter 2 numbers")
////
////		var a, b int
////
////		fmt.Scan(&a, &b)
////		ch1 <- a
////		ch2 <- b
////	}()
////	//	time.Sleep(7*time.Second)
////	var operator string
////	fmt.Println("enter operator")
////	fmt.Scanln(&operator)
////	if operator == "+" {
////		sum(ch1, ch2)
////	} else if operator == "-" {
////		sub(ch1, ch2)
////	} else if operator == "*" {
////		mul(ch1, ch2)
////	} else {
////		div(ch1, ch2)
////	}
////}
////
////package main
////
////import "fmt"
////
////func main() {
////	for i := 0; i < 6; i++ {
////		for j := 0; j < 6; j++ {
////			if i == 0 && j == 0 || i == 0 && j == 5 || i == 1 && j == 4 || i == 1 && j == 1 || i == 4 && j == 1 || i == 4 && j == 4 || i == 5 && j == 0 || i == 5 && j == 5 ||i=={
////				fmt.Print("H")
////
////			} else {
////				fmt.Print(" ")
////
////			}
////
////		}
////		fmt.Println(" ")
////*/
////*/
////	}
////}
////
////
