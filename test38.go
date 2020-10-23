package main
import (
      "fmt"
      "runtime"
      "sync"
     )

func main(){
    runtime.GOMAXPROCS(4)
    var counter uint64
    var wg sync.WaitGroup

    for i:=0;i<50;i++{
      wg.Add(1)
      go func(){
        defer wg.Done()
        for c:=0;c<1000;c++{
          counter++
        }
      }()
    }
    wg.Wait()
    fmt.Printf("Counter:",counter)
}