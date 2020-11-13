package main

import (
    "fmt"
    "reflect"
    "regexp"

    "gopkg.in/go-playground/validator.v8"
)

// Person ...
type Person struct {
    ID       string `json:"id" validate:"required,min=36,max=36"`
    Password string `json:"id" validate:"upwd"`
}

var (
    validate       = validator.New(&validator.Config{TagName: "validate"})
    passwordString = "^[a-zA-Z0-9]$" // regex that compiles
    passwordRegex  = regexp.MustCompile(passwordString)
)

func main() {

    validate.RegisterValidation("validpasswd", func(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
        return passwordRegex.MatchString(field.String())
    })
    validate.RegisterAliasValidation("upwd", "min=8,max=20,validpasswd")

    var errMsg string

    person := Person{
        ID:       "36",
        Password: "!!!!!!!!!!!!!!",
    }

    if errs := validate.Struct(person); errs != nil {

        if err := errs.(validator.ValidationErrors)["Person.Password"]; err != nil {
            if err.Tag == "upwd" {
                errMsg = "invalid password!"
            }
        }
    }

    fmt.Println(errMsg)
}
PS C:\Users\Sudha\Desktop\golang\sudha\assign\passwordmatching> go run pswd1.go;


prgram:
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str1 := "this is a [sample] [[string]] with [SOME] special words"

	re := regexp.MustCompile(`\[([^\[\]]*)\]`)
	fmt.Printf("Pattern: %v\n", re.String())      // print pattern
	fmt.Println("Matched:", re.MatchString(str1)) // true

	fmt.Println("\nText between square brackets:")
	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		element = strings.Trim(element, "[")
		element = strings.Trim(element, "]")
		fmt.Println(element)
	}
}

output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\passwordmatching> go run pswd2.go;
Pattern: \[([^\[\]]*)\]
Matched: true

Text between square brackets:
sample
string
SOME

program:

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "We @@@Love@@@@ #Go!$! ****Programming****Language^^^"

	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1))        // true

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}

output:


PS C:\Users\Sudha\Desktop\golang\sudha\assign\passwordmatching> go run pswd3.go;
Pattern: [^a-zA-Z0-9]+
true
 @@@
@@@@ #
!$! ****
****
^^^

program:

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "If I am 20 years 10 months and 14 days old as of August 17,2016 then my DOB would be 1995-10-03"

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern

	fmt.Println(re.MatchString(str1)) // true

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\passwordmatching> go run pswd4.go;
Pattern: \d{4}-\d{2}-\d{2}
true
1995-10-03

program:

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := `Proxy Port Last Check Proxy Speed Proxy Country Anonymity 118.99.81.204
	118.99.81.204 8080 34 sec Indonesia - Tangerang Transparent 2.184.31.2 8080 58 sec 
	Iran Transparent 93.126.11.189 8080 1 min Iran - Esfahan Transparent 202.118.236.130 
	7777 1 min China - Harbin Transparent 62.201.207.9 8080 1 min Iraq Transparent`

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1)) // true

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}

output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\passwordmatching> go run pswd5.go;
Pattern: (25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}
true
118.99.81.204
118.99.81.204
2.184.31.2
93.126.11.189
202.118.236.130
62.201.207.9


program::
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := `http://www.suon.co.uk/product/1/7/3/`

	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1)) // true

	submatchall := re.FindAllString(str1,-1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}

output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\passwordmatching> go run pswd6.go;
Pattern: ^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)
true
http://www.suon.co.uk

program::
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "ç$€§/az@gmail.com"
	str2 := "abcd@gmail_yahoo.com"
	str3 := "abcd@gmail-yahoo.com"
	str4 := "abcd@gmailyahoo"
	str5 := "abcd@gmail.yahoo"

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern	
	fmt.Printf("\nEmail: %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("Email: %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("Email: %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("Email: %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("Email: %v :%v\n", str5, re.MatchString(str5))
}

output:
PS C:\Users\Sudha\Desktop\golang\sudha\assign\passwordmatching> go run pswd7.go;
Pattern: ^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$

Email: ç$€§/az@gmail.com :false
Email: abcd@gmail_yahoo.com :false
Email: abcd@gmail-yahoo.com :true
Email: abcd@gmailyahoo :true
Email: abcd@gmail.yahoo :true
