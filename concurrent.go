package main

import "fmt"

func goAdd(x, y int, ch chan int) {
	ch <- 1
	fmt.Println("hello world1")
}
func main() {
	var ch chan int
	ch=make(chan int,1)
	go goAdd(1, 2, ch)
	<-ch

}
