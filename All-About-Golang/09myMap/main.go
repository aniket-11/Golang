package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Here is Maps AWWW ")
	languages :=  make(map[string]string)

	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["JS"] = "JavaScript"
	languages["CPP"] = "C++"

	fmt.Println("All languages are ",languages)

	fmt.Println("JS shorts of an ",languages["JS"])

	//to remove and element is map we use delete in it

	delete(languages,"RB")

	fmt.Println("Here is New list = ",languages)

	//looping in go similar tpo foreach in javaScript

	for key,value := range languages {
		fmt.Printf("The Key %v have an value is %v\n",key,value)
	}

	//By using an comma okay syntax

	for value := range languages {
		fmt.Printf("The value is = %v\n",languages[value])
	}
}