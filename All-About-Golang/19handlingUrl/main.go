package main 

import (
	"fmt"
	"net/url"
	//"io/ioutil"
)
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj516knvkjnf"
func main()  {
	fmt.Println("Here is handling URL")
	fmt.Println(myurl)

	//parsing an url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query programs are : %T \n",qparams)

	fmt.Println("courseName = ",qparams["coursename"])

	for _,val := range qparams {
		fmt.Println("Param is : ",val)
	}

	partsOfUrl := &url.URL {
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user-aniket",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}