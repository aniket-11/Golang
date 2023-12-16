package main
import (
	"fmt"
)
type User struct {
	Name string
	Email string
	Status bool 
}

func main()  {
	aniket := User{"Aniket","aniket@dev.in",true}
	
	aniket.getStatus()
	aniket.newEmail()

	fmt.Println("User Email =",aniket.Email)
}

func (u User) getStatus()  {
	fmt.Println("User status =",u.Status)
}

func (u *User) newEmail()  {
	u.Email = "test@dev.go"
	fmt.Println("New Email =",u.Email)
}