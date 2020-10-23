package main
import "fmt"

func main(){
	var arr[5]int=[5]int{1,2,3,4,5}
	for i:=0 ;i<5;i++{
		fmt.Println(arr[i])
	}
	fmt.Println("Displaying second array")
	brr:=[6]int{1,2,3,4,5,6}
	for i:=0;i<6;i++{
		fmt.Println(brr[i])
	}
}