program:

package main

import (
    "fmt"
    "time"
)

// The pinger prints a ping and waits for a pong
func pinger(pinger <-chan int, ponger chan<- int) {
    for {
        <-pinger
        fmt.Println("ping")
        time.Sleep(time.Second)
        ponger <- 1
    }
}

// The ponger prints a pong and waits for a ping
func ponger(pinger chan<- int, ponger <-chan int) {
    for {
        <-ponger
        fmt.Println("pong")
        time.Sleep(time.Second)
        pinger <- 1
    }
}

func main() {
    ping := make(chan int)
    pong := make(chan int)

    go pinger(ping, pong)
    go ponger(ping, pong)

    // The main goroutine starts the ping/pong by sending into the ping channel
    ping <- 1

    for {
        // Block the main thread until an interrupt
        time.Sleep(  10 * time.Second)
        break
    }
}
output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\pingpong> go run ping4.go;
ping
pong
ping
pong
ping
pong
ping
pong
ping
pong

program:

package main
 
import (
	"fmt"
	"time"
)
 
func main() {
	pingChan := make(chan int, 1)
	pongChan := make(chan int, 1)
 
	go ping(pingChan, pongChan)
	go pong(pongChan, pingChan)
 
	pingChan <- 1
 
	select {}
}
 
func ping(pingChan <-chan int, pongChan chan<- int) {
	for {
		ball := <-pingChan
 
		fmt.Println("Ping", ball)
		time.Sleep(1 * time.Second)
 
		pongChan <- ball+1
	}
}
 
func pong(pongChan <-chan int, pingChan chan<- int) {
	for {
		ball := <-pongChan
 
		fmt.Println("Pong", ball)
		time.Sleep(1 * time.Second)
 
		pingChan <- ball+1

		for {
			// Block the main thread until an interrupt
			time.Sleep(time.Second)
		}
	}
}

    
output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\pingpong> go run ping3.go;
Ping 1
Pong 2
Ping 3
Pong 4
Ping 5
Pong 6
Ping 7
Pong 8
Ping 9
Pong 10

program:

package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan struct{}, 1)
	pong := make(chan struct{}, 1)

	ping<- struct{}{}

	go play(ping, pong)

	time.Sleep(time.Millisecond)
}

func play(ping, pong chan struct{}) {
	for {
		select {
		case <-ping:
			fmt.Println("ping")
			pong<- struct{}{}
		case <-pong:
			fmt.Println("    pong")
			ping<- struct{}{}
		}
	}
}

output:

PS C:\Users\Sudha\Desktop\golang\sudha\assign\pingpong> go run ping2.go;
ping
pong
ping
    pong
ping
    pong
ping
    pong
ping
    pong
ping
    pong
ping
    pong
ping


program:

package main


import "fmt"


func main() {
	
pingChan := make(chan string)
	
pongChan := make(chan string)

	
go printer(pongChan)
	
go pinger(pingChan)
	
go ponger(pingChan, pongChan)

	


var input string
	
fmt.Scanln(&input)

}



func pinger(pingChan chan<- string) {

	
fmt.Println("pinger sending \"ping\"")
	pingChan <- "ping"

}



func ponger(pingChan <-chan string, pongChan chan<- string) {
	
fmt.Println("ponger received", <-pingChan)

	// respond with a pong
	

fmt.Println("ponger replying with \"pong\"")
	pongChan <- "pong"

}



func printer(pongChan <-chan string) {

fmt.Println("printer received", <-pongChan)

}
output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\pingpong> go run ping.go;
pinger sending "ping"
ponger received ping
ponger replying with "pong"
printer received pong
f
