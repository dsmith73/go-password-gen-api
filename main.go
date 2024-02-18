// api service to generate a password and return it in json

package main

import (
	// 	"encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Pass struct is used to format the password string for the json response
type Pass struct {
	Password string `json:"password"`
}

type Arguments struct {
	PassLength int `json:"length"`
	Lower bool `json:"lowercase"`
	Upper bool `json:"uppercase"`
	Num bool `json:"number"`
	Sym bool `json:"symbol"`
	FirstChar int `json:"firstChar"`
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	fmt.Println("password started on localhost:8088")
	myRouter.HandleFunc("/", homePage)
	// POST vars to endpoint to create password
	myRouter.HandleFunc("/generate", generate).Methods("POST") 
	// GET password with modified values
	myRouter.HandleFunc("/generate/{passLength}/{lower}/{upper}/{num}/{sym}/{firstChar}", generate).Methods("GET")
	// GET password from endpoint with default values
	myRouter.HandleFunc("/generate", generate).Methods("GET") //  genPWD(11,true,true,true,true,3)


	
	log.Fatal(http.ListenAndServe(":8088", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {

    // fmt.Println(r.URL.Path)
    p := "." + r.URL.Path
    if p == "./" {
        p = "./public/index.html"
    }
    http.ServeFile(w, r, p)

}

func generate(w http.ResponseWriter, r *http.Request) {
	var password string

	if r.Body != nil && r.Method == "POST" {
		var arguments Arguments
		arguments.PassLength = 11
		arguments.Lower = true
		arguments.Upper = true
		arguments.Num = true
		arguments.Sym = false
		arguments.FirstChar = 3

		json.NewDecoder(r.Body).Decode(&arguments)

		fmt.Println("POST: /generate")


		password = genPWD(arguments.PassLength, arguments.Lower, arguments.Upper, arguments.Num, arguments.Sym, arguments.FirstChar)

		

		
	} else {
		fmt.Println("GET: /generate")

		vars := mux.Vars(r)
		passLength, err := strconv.Atoi(vars["passLength"])
		lower, err := strconv.ParseBool(vars["lower"])
		upper, err := strconv.ParseBool(vars["upper"])
		num, err := strconv.ParseBool(vars["num"])
		sym, err := strconv.ParseBool(vars["sym"])
		firstChar, err := strconv.Atoi(vars["firstChar"])
	
		if err != nil {
			// log.Fatal(err)
			passLength = 13
			lower = true
			upper = true
			num = true
			sym = true
			firstChar = 3
			password = genPWD(passLength, lower, upper, num, sym, firstChar)
		} else {
	
			password = genPWD(passLength, lower, upper, num, sym, firstChar)
		}
	

	}


	pwd := Pass{
		Password: password,
	}

	w.Header().Set("Content-Type", "application/json")
    // w.WriteHeader(http.StatusCreated) 	// status 201  
    w.WriteHeader(http.StatusOK) 	// status 200

	json.NewEncoder(w).Encode(pwd)
	// // fmt.Println(string(pwd))

}

func genPWD(passLength int, lower bool, upper bool, num bool, sym bool, firstChar int) string {
	symbols := ".?],{+*:[$#@}=!"
	numbers := "6781235904"
	upCase := "OPWXKQRAFGHIJMNSCDTUBZEVLY"
	loCase := "awxbfhirscdjzklmnopqetuvyg"

	fmt.Println("Generating password with:",
		"\nLength:    ", passLength,
		"\nNumbers:   ", num,
		"\nSymbols:   ", sym,
		"\nLowercase: ", lower,
		"\nUppercase: ", upper)

	// seed rand with time to ensure random output
	// without the seed, you always get the same output, because seed defaults to 1
	rand.Seed(time.Now().UnixNano())

	// var password string
	fl := ""
	pl := ""
	password := ""

	// set possible character types for the first character of the password
	if firstChar == 1 && lower == true {
		fl = loCase
	}
	if firstChar == 2 && upper == true {
		fl = upCase
	}
	if firstChar == 3 && lower == true && upper == true {
		fl = loCase + upCase
	}
	if firstChar == 4 && num == true {
		fl = numbers
	}

	// add possible characters to password string
	if lower == true {
		pl = pl + loCase
	}

	if num == true {
		pl = pl + numbers
	}

	if upper == true {
		pl = pl + upCase
	}

	if sym == true {
		pl = symbols + pl
	}

	if firstChar == 5 {
		fl = pl

		if num == false || sym == false || lower == false || upper == false {
			fmt.Println("Received the following switches:\nNumbers:   ", num,
				"\nSymbols:   ", sym,
				"\nLowercase: ", lower,
				"\nUppercase: ", upper,
				"\nProgressing with the switches which were available...")
		}
	}

	if passLength < 6 || passLength > 100 {
		fmt.Println("ERROR: You must pick a length between 6 & 100 characters.")
	}

	// generate the first character
	password = string(fl[rand.Intn(len(fl))])

	// generate and append the rest of the password
	for i := 0; i < passLength-1; i++ {
		password = password + string(pl[rand.Intn(len([]rune(pl)))])
	}

	return password
}

func main() {

	// passLength := 13
	// lower := true
	// upper := true
	// num := false
	// sym := true
	// firstChar := 5
	
	// genPWD(passLength, lower, upper, num, sym, firstChar)

	handleRequests()
}


