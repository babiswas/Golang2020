package main
import (
   "fmt"
   "os"
)

func main(){
   for i:=1;i<len(os.Args);i++{
       fmt.Println(os.Args[i])
   }
   fmt.Println("Looping Completed")
   fmt.Println(os.Args)
}