package main
 
import(
	"fmt"
)

type User struct {
	Name string
	Email string
	Status bool
	Age int

}

func main()  {
	aniket := User {"Aniket","aniket@go.dev",true,24}

	malli :=  User{"Malli","Malli@go.dev",true,22}

	fmt.Println(aniket)

	fmt.Printf("Aniket Details are = %+v",aniket)

	fmt.Printf("Name = %v, Age = %v, Status = %v, Email = %v",aniket.Name,aniket.Age,aniket.Status,aniket.Email)

	fmt.Printf("Name = %v, Age = %v, Status = %v, Email = %v",malli.Name,malli.Age,malli.Status,malli.Email)
}