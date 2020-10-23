package main
import "fmt"

func main(){
   arr:=make([]int,5)
   for i:=0;i<len(arr);i++{
        arr[i]=10+i
   }
   for i:=0;i<len(arr);i++{
       fmt.Println(arr[i])
   }
   arr=append(arr,2,3,4)
   fmt.Println(arr)
   fmt.Println(arr[1:4])
}