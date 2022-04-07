package main
import "fmt"

func channel() {
	dataChan := make(chan int) // channell instantiated and saved in variable dataChan
	go func() {                // a function to pass a value into a channel and return the value into a variable
		dataChan <- 25
	}()
	n := <-dataChan
	fmt.Printf("n = %v\n", n)
}
 