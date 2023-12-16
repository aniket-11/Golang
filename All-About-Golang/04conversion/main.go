package main 
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)
func main()  {
	fmt.Println("Enter rating for our pizza app")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("The voting are, ",input)
	fmt.Printf("The type of input is %T",input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err!= nil {
		fmt.Println(err)
	}else { 
		fmt.Println("Added 1 to your rating ",numRating+1 )
	}

}