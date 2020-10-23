package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1,_:=strconv.ParseInt(scanner.Text(),10,64)
	scanner.Scan()
	input2,_:=strconv.ParseInt(scanner.Text(),10,64)
	sum:=input1+input2
	fmt.Println(sum)
}