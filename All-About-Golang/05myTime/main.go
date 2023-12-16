package main
import (
	"time"
	"fmt"
)
func main()  {
	curentTime := time.Now()
	fmt.Println(curentTime.Format("01-02-2006 15:04:05 Monday"))
	createDate := time.Date(2020, time.August, 10,23,23,0,0,time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-06 Monday"))
}