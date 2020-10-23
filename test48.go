package main
import (
    "fmt"
    "sync"
    "time"
   )

func main(){
   var wg sync.WaitGroup
   wg.Add(1)
   go func(){
      defer wg.Done()
      fmt.Println("Obtained task")
      time.Sleep(10*time.Millisecond)
      fmt.Println("Completed task")
   }()

   wg.Add(1)

   go func(){
     defer wg.Done()
     fmt.Println("Obtained task1")
     time.Sleep(10*time.Millisecond)
     fmt.Println("Completed task")
   }()

   wg.Wait()   
}