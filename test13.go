package main
import "fmt"

func main(){
	arr:=make([]int,6)
	for i:=0;i<len(arr);i++{
		arr[i]=10+i
	}
	s:=make([]int,6)
	copy(s,arr)
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

}