package main
import (
        "fmt"
        "runtime"
        "sync"
        "atomic"
       )

func main(){
   runtime.GOMAXPROCS(4)
   var counter uint64
   var wg sync.WaitGroup

   for i:=0;i<50;i++{
     wg.Add(1)
     go func(){
       defer wg.Done()
       for c:=0;c<100;c++{
         atomic.AddUint64(&counter,1)
       }
     }()
   }

   wg.Wait()
   fmt.Println(counter)
}