package main
import "fmt"
import "os"
import "bufio"
func main()  {
	welcome := "Welcome to user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating")
	input, _ := reader.ReadString('\n')
	fmt.Println("thnks for voting",input)
	fmt.Printf("the type of input is %T",input)

}