package main 
import (
	"fmt"
	"os"
	"io/ioutil"
	"io"
)

func main()  {
	fmt.Println("Yey Here is an file Handling in Golang")
	content := "This need to go in file - codeaniket@dev.go"

	//create a file
	file, err := os.Create("./myFile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file,content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is : ",length)
	defer file.Close()
	readFile("./myFile.txt")

}


func readFile(filename string)  {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inide file is \n",string(databyte))
}