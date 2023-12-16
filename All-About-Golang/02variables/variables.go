package main 
import "fmt"
const VariableToken string = "hbschivb" //public if the variable name start with capital ina global
func main()  {
	var userName string = "Aniket"
	var salaryAniket float32 = 75000.3542
	var inputNumber int = 355

	var averageSalary float64

	fmt.Println(userName)
	fmt.Printf("The Variable Type : %T \n",userName)
	fmt.Println(salaryAniket)
	fmt.Printf("The Variable Type : %T \n",salaryAniket)
	fmt.Println(inputNumber)
	fmt.Printf("The Variable Type : %T \n",inputNumber)

	averageSalary = 35555.00
	taxPayble := 35000

	fmt.Println(averageSalary)
	fmt.Printf("The Variable Type : %T \n",averageSalary)

	fmt.Println(taxPayble)
	fmt.Printf("The Variable Type : %T \n",taxPayble)

	fmt.Println(VariableToken)
	fmt.Printf("The Variable Type : %T \n",VariableToken)
}