package main

import "fmt"

func main() {
	c := make(chan int)
	q := 10
	go generateInts(c, q)

	for i := range c {
		fmt.Println(i)
	}

}

func generateInts(c chan int, q int) {
	for i := 0; i < q; i++ {
		c <- i
	}
	close(c)
}

program:
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	q := 10
	go worker(c, 1)
	go worker(c, 2)
	go worker(c, 3)
	go generateInts(c, q)

	time.Sleep(10 * time.Second)
}

func worker(c chan int, id int) {
	for i := range c {
		fmt.Printf("Worker %d received value %d\n", id, i)
		time.Sleep(1 * time.Second)
	}
}

func generateInts(c chan int, q int) {
	for i := 0; i < q; i++ {
		c <- i
		fmt.Printf("Sent value %d\n", i)
		time.Sleep(1 * time.Millisecond)
	}
	close(c)
}

program:
type data struct {
	// ...
}

func handleHTTPRequest(w http.ResponseWriter, r *http.Request) {
	// ...
	c := make(chan data)
	defer close(c)
	defer r.Body.Close()
	go saveDataFromChannel(c)
	c <- calculateStuff(getData(r.Body))
}

func getData(r io.Reader) data {
	// ...
	return data{}
}

func calculateStuff(d data) data {
	// ...
	return d
}

func saveDataFromChannel(c chan data) {
	for d := range c {
		saveToDatabase(d)
		// ...
	}
}

func saveToDatabase(d data) {
	// ...
}