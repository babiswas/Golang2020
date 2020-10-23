package main
import (
       "fmt"
       "time"
     )


func fun(s string){
   for i:=0;i<3;i++{
      fmt.Println(s)
      time.Sleep(1*time.Millisecond)
   }
}


func main(){
   fun("direct call")
   go fun("go routine")
   go func(){
      fun("Hello Hello")
   }()
   time.Sleep(100*time.Millisecond)
   fmt.Println("Done")
}