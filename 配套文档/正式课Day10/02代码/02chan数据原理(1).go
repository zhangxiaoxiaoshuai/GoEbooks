package main
import (
	"fmt"
	"time"
	"unsafe"
)
func goRoutineA(a <-chan int) {
	val := <-a
	fmt.Println("goRoutineA:", val)
}
func goRoutineB(b chan int) {
	val := <-b
	fmt.Println("goRoutineB:", val)
}
func main() {
	ch := make(chan int, 100)
	go goRoutineA(ch)
	go goRoutineB(ch)
	ch <- 3
	time.Sleep(time.Second)
	fmt.Println(unsafe.Sizeof(ch))
}