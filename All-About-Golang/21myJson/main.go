package main

import (
	"fmt"
	"encoding/json"
)




func main()  {
	fmt.Println("Here is an creation of json structure ")
	EncodeJson()

}

func EncodeJson(){type course struct {
	Name string
	Price int
	Platform string
	Password string
	Tags []string
}

func main()  {
	fmt.Println("Here is an creation of json structure ")
	EncodeJson()

}

func EncodeJson(){
	lcoCourses := []course {
		{"ReactJs Bootcamp",299,"LearnCodeOnline.in","abc123",[]string{"web-dev","js"}},
		{"Mern Bootcamp",299,"LearnCodeOnline.in","def123",[]string{"full-dev","js"}},
		{"Angular Bootcamp",299,"LearnCodeOnline.in","hdfc254",nil},
	}

	finalJson, err:= json.Marshal(lcoCourses)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n",finalJson)
}



// Above code is having some bug ! Do not use it.
