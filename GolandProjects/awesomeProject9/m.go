package main

//
//import (
//	"fmt"
//	"os"
//)
//
//type TV interface {
//	on()
//	off()
//}
//
//type s1 struct{}
//
//func (o s1) on() {
//	fmt.Println("ON")
//}
//
//func (o s1) off() {
//	fmt.Println("ON")
//}
//
//type INV interface {
//	execute()
//}
//
//type on struct {
//	o s1
//}
//
//type off struct {
//	o s1
//}
//
//func (o on) execute() {
//	o.o.on()
//}
//
//func (o off) execute() {
//	o.o.off()
//}
//
//type s2 struct {
//	o INV
//}
//
//func (o s2) press() {
//	o.o.execute()
//}
//
//func main() {
//
//	os.Setenv("1", "vat")
//	fmt.Println("HELLO", os.Getenv("1"))
//
//}
