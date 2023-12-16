package main
import(
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"net/url"
)

func main()  {
	fmt.Println("Here is an handling requests")
	//PerformgetRequest()
	//PerformPostJsonRequest()
	PerformFormreuest()
}


//to performing an get request

func PerformPostJsonRequest(){
	const myUrl = "http://localhost:3000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform": "learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myUrl,"application/json",requestBody)
	
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}



//performing get request

func PerformgetRequest()  {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl) 
		if err !=  nil {
			panic(err)
		}


	defer response.Body.Close()	

	fmt.Println("Status code: ",response.StatusCode)

	fmt.Println("Content Length: ",response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is  :",byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(content)
	// fmt.Println(string(content))
		
	
}

func PerformFormreuest()  {
	const myUrl = "http://localhost:3000/postform" 

	//fake json paylod formd data 
	data := url.Values{}
	data.Add("firstName","Aniket")
	data.Add("lastName","Kshirsagar")
	data.Add("Email","Aniket@dev.go")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
	
}