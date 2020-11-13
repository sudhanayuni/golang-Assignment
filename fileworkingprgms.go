package main
// import the 2 modules we need
import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("localfile.data")
    
    if err != nil {
        fmt.Println(err)
    }
    fmt.Print(string(data))

}


program::

package main

import (
    "fmt"
    "io/ioutil"
)

func main() {

    mydata := []byte("All the data I wish to write to a file")

    // the WriteFile method returns an error if unsuccessful
    err := ioutil.WriteFile("myfile.data", mydata, 0777)
    // handle this error
    if err != nil {
        // print it out
        fmt.Println(err)
    }

    data, err := ioutil.ReadFile("myfile.data")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

}

program:

package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {

    mydata := []byte("All the data I wish to write to a file\n")

    // the WriteFile method returns an error if unsuccessful
    err := ioutil.WriteFile("myfile.data", mydata, 0777)
    // handle this error
    if err != nil {
        // print it out
        fmt.Println(err)
    }

    data, err := ioutil.ReadFile("myfile.data")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

    f, err := os.OpenFile("myfile.data", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
        panic(err)
    }

    data, err = ioutil.ReadFile("myfile.data")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

}

program:
package main
 
import (
	"bufio"
	"fmt"
	"log"
	"os"
)
 
func main() {
	file, err := os.Open("test.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}
program:
package main
 
import (
	"log"
	"os"
)
 
func main() {
	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}
program:
package main
 
import (
	"log"
	"os"
)
 
func main() {
	_, err := os.Stat("test")
 
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
 
	}
}