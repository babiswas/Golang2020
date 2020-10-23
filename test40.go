package main
import (
      "fmt"
      "sync"
      "time"
    )

var src=make(map[string]string)

func main(){
   var wg sync.WaitGroup
   mu:=sync.Mutex{}
   c:=sync.NewCond(&mu)
   wg.Add(1)
   go func(){
     defer wg.Done()
     c.L.Lock()
     for len(src)==0{
        time.Sleep(1*time.Millisecond)
        c.Wait()
     }
     fmt.Println(src["rsc1"])
     c.L.Unlock()
   }()

   c.L.Lock()
   src["rsc1"]="foooo"
   c.Signal()
   c.L.Unlock()
   wg.Wait()   
}