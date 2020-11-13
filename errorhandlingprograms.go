program:

package main
import "fmt"
import "os"
import "errors"

//function accepts a filename and tries to open it.
func fileopen(name string) (string, error) {
    f, er := os.Open(name)

    //er will be nil if the file exists else it returns an error object  
    if er != nil {
        //created a new error object and returns it  
        return "", errors.New("Custom error message: File name is wrong")
    }else{
    	return f.Name(),nil
    }
}

func main() {  
    //receives custom error or nil after trying to open the file
    filename, error := fileopen("invalid.txt")
    if error != nil {
        fmt.Println(error)
    }else{
    	fmt.Println("file opened", filename)
    }  
}

output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\errorhandling> go run errorhandling.go;
Custom error message: File name is wrong

program:

package main
import "fmt"
import "os"

//function accepts a filename and tries to open it.
func fileopen(name string) {
    f, er := os.Open(name)

    //er will be nil if the file exists else it returns an error object  
    if er != nil {
        fmt.Println(er)
        return
    }else{
    	fmt.Println("file opened", f.Name())
    }
}

func main() {  
    fileopen("invalid.txt")
}


output:

PS C:\Users\Sudha\Desktop\golang\sudha\assign\errorhandling> go run errorhandling2.go;
open invalid.txt: The system cannot find the file specified.


program:

package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Open("/test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}

output:

PS C:\Users\Sudha\Desktop\golang\sudha\assign\errorhandling> go run errorhandling3.go;
open /test.txt: The system cannot find the file specified.

program:
package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Open("test.txt")
    if err != nil {
        if pErr, ok := err.(*os.PathError); ok {
            fmt.Println("Failed to open file at path", pErr.Path)
            return
        }
        fmt.Println("Generic error", err)
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}

output:

PS C:\Users\Sudha\Desktop\golang\sudha\assign\errorhandling> go run errorhandling4.go;
Failed to open file at path test.txt

program:

package main

import (  
    "fmt"
    "net"
)

func main() {  
    addr, err := net.LookupHost("golangbot123.com")
    if err != nil {
        if dnsErr, ok := err.(*net.DNSError); ok {
            if dnsErr.Timeout() {
                fmt.Println("operation timed out")
                return
            }
            if dnsErr.Temporary() {
                fmt.Println("temporary error")
                return
            }
            fmt.Println("Generic DNS error", err)
            return
        }
        fmt.Println("Generic error", err)
        return
    }
    fmt.Println(addr)
}


output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\errorhandling> go run errorhandling5.go;
Generic DNS error lookup golangbot123.com: no such host

program:

package main

import (  
    "fmt"
    "path/filepath"
)

func main() {  
    files, err := filepath.Glob("[")
    if err != nil {
        if err == filepath.ErrBadPattern {
            fmt.Println("Bad pattern error:", err)
            return
        }
        fmt.Println("Generic error:", err)
        return
    }
    fmt.Println("matched files", files)
}

output:

PS C:\Users\Sudha\Desktop\golang\sudha\assign\errorhandling> go run errorhandling6.go;
Bad pattern error: syntax error in pattern

package main

import (  
    "fmt"
    "path/filepath"
)

func main() {  
    files, _ := filepath.Glob("[")
    fmt.Println("matched files", files)
}

output:

PS C:\Users\Sudha\Desktop\golang\sudha\assign\errorhandling> go run errorhandling7.go;
matched files []


