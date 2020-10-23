package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the first number")
	scanner.Scan()
	a,_:=strconv.ParseInt(scanner.Text(),10,64)
	fmt.Println("Enter the second number")
	scanner.Scan()
	b,_:=strconv.ParseInt(scanner.Text(),10,64)
	var sum int64=a+b
	fmt.Println("The sum of the number is ",sum)
	
}