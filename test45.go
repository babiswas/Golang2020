package main
import (
    "fmt"
    "bufio"
	"os"
	"strconv"
	"math"
)

func find_hypotenuse()int64{
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the value of perpendicular")
	scanner.Scan()
	x,_:=strconv.ParseInt(scanner.Text(),10,64)
	fmt.Println("Enter the value of base")
	scanner.Scan()
	y,_:=strconv.ParseInt(scanner.Text(),10,64)
	return x*x+y*y
}

func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the value of base")
	scanner.Scan()
	x,_:=strconv.ParseInt(scanner.Text(),10,64)
	fmt.Println("enter the value of perpendicular")
	scanner.Scan()
	y,_:=strconv.ParseInt(scanner.Text(),10,64)
	fmt.Println("The value of the hypotenuse is:")
	sum_square:=x*x+y*y
	fmt.Println("The square of the hypotenuse is:",sum_square)
	fmt.Println(math.Sqrt(float64(sum_square)))
	fmt.Println("the hypotenuse is",find_hypotenuse())
	fmt.Println("The hypotenuse is",math.Sqrt(float64(find_hypotenuse())))	
}