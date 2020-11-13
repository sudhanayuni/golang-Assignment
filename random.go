package main 
  
import ( 
    "fmt"
    "math/rand"
    "time"
) 
  
func main() { 
  
    fmt.Println(rand.Intn(500)) 
    fmt.Println(rand.Intn(500)) 
    fmt.Println(rand.Intn(500)) 
    fmt.Println() 
  
    fmt.Println(rand.Float64()) 
  
    fmt.Println((rand.Float64() * 8) + 7) 
    fmt.Println((rand.Float64() * 8) + 7) 
    fmt.Println() 
  
    
    x1 := rand.NewSource(time.Now().UnixNano()) 
    y1 := rand.New(x1) 
      
    fmt.Println(y1.Intn(500)) 
    fmt.Println(y1.Intn(500)) 
    fmt.Println() 
  
    x2 := rand.NewSource(55) 
    y2 := rand.New(x2) 
    fmt.Println(y2.Intn(200)) 
    fmt.Println(y2.Intn(200)) 
    fmt.Println() 
      
    x3 := rand.NewSource(7) 
    y3 := rand.New(x3) 
    fmt.Println(y3.Intn(200)) 
    fmt.Println(y3.Intn(200)) 
} 