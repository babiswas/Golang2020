package main
import (
        "fmt"
        "sync"
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
      for len(src)<1{
         c.Wait()
      }
     fmt.Println(src["rsc1"])
     c.L.Unlock()
   }()
   wg.Add(1)
   go func(){
      defer wg.Done()
      c.L.Lock()
      for len(src)<2{
         c.Wait()
      }
      fmt.Println(src["rsc2"])
      c.L.Unlock()
   }()
   c.L.Lock()
   src["rsc1"]="value1"
   src["rsc2"]="value2"
   c.Broadcast()
   c.L.Unlock()
   wg.Wait()   
}