package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)
const url = "https://lco.dev"

func main()  {
	fmt.Println("Here is an handling web requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response if of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil{
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}