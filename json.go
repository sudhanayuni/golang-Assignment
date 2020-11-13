package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    // define data structure 
    type DayPrice struct {
      Gold int
      Silver int
      Platinum int
    }

    // json data
    data := `{"gold": 1271,"silver": 1284,"platinum": 1270}`        
    var obj DayPrice

    // unmarshall it
    err := json.Unmarshal([]byte(data), &obj)
    if err != nil {
        fmt.Println("error:", err)
    }

    // can access using struct now
    fmt.Printf("Gold     : $ %d\n", obj.Gold);
    fmt.Printf("Silver   : $ %d\n", obj.Silver);
    fmt.Printf("Platinum : $ %d\n", obj.Platinum);
}

program:
package main

import "os"
import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

// Define data structure 
type Response struct {
  TradeID    int
  Price      string
  Size       string
  Bid        string
  Ask        string
  Volume     string
  Time       string
}


func get_content() {
  // json data
  url := "https://api.gdax.com/products/BTC-EUR/ticker"

  res, err := http.Get(url)

  if err != nil {
       panic(err.Error())
     }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
         panic(err.Error())
       }

      var data Response

      // unmarshall
      json.Unmarshal(body, &data)
      //fmt.Printf("Results: %v\n", data)

      // print values of the object
      fmt.Printf("Price: $ %s\n", data.Price)      
      fmt.Printf("Price: $ %s\n", data.Bid)
      fmt.Printf("Price: $ %s\n", data.Ask)

      os.Exit(0)
}

func main() {
  get_content()
}

program:
package main

import (
	"encoding/json"
	"fmt"
)

// The same json tags will be used to encode data into JSON
type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func main() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// we can use the json.Marhal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))
}
program:
package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string `json:"birdType"`
	// we can set the "omitempty" property as part of the JSON tag
	Description string `json:"what it does,omitempty"`
}

func main() {
	pigeon := &Bird{
		Species:     "Pigeon",
	}

	data, _ := json.Marshal(pigeon)

	fmt.Println(string(data))
}

program:
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// The keys need to be strings, the values can be
	// any serializable value
	birdData := map[string]interface{}{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	// JSON encoding is done the same way as before	
	data, _ := json.Marshal(birdData)
	fmt.Println(string(data))
}
program:
package main 
   
import ( 
    "fmt"
    "encoding/json"
) 
   
// declaring a struct 
type Human struct{ 
       
    // defining struct variables 
    Name string 
    Address string 
    Age int
} 
   
// main function 
func main() { 
       
    // defining a struct instance 
    var human1 Human 
       
    // data in JSON format which 
    // is to be decoded 
    Data := []byte(`{ 
        "Name": "Deeksha",   
        "Address": "Hyderabad", 
        "Age": 21 
    }`) 
       
    // decoding human1 struct 
    // from json format 
    err := json.Unmarshal(Data, &human1) 
       
    if err != nil { 
           
        // if error is not nil 
        // print error 
            fmt.Println(err) 
    } 
       
    // printing details of 
    // decoded data  
    fmt.Println("Struct is:", human1) 
    fmt.Printf("%s lives in %s.\n", human1.Name, human1.Address) 
       
    // unmarshaling a JSON array 
    // to array type in Golang 
       
    // defining an array instance 
    // of struct type 
    var human2 []Human 
       
    // JSON array to be decoded 
    // to an array 
    Data2 := []byte(` 
    [ 
        {"Name": "Vani", "Address": "Delhi", "Age": 21}, 
        {"Name": "Rashi", "Address": "Noida", "Age": 24}, 
        {"Name": "Rohit", "Address": "Pune", "Age": 25} 
    ]`) 
       
    // decoding JSON array to  
    // human2 array 
    err2 := json.Unmarshal(Data2, &human2) 
       
        if err2 != nil { 
       
        // if error is not nil 
        // print error 
            fmt.Println(err2) 
        } 
       
    // printing decoded array  
    // values one by one 
    for i := range human2{ 
       
        fmt.Println(human2[i]) 
    } 
} 