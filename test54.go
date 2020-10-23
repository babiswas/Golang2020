package main
import (
   "fmt"
   "sync"
   "time"
)

func task1(wg *sync.WaitGroup){
   defer wg.Done()
   fmt.Println("Task1 is obtained")
   time.Sleep(10*time.Millisecond)
   fmt.Println("Task2 completed")
}

func task2(wg *sync.WaitGroup){
   defer wg.Done()
   fmt.Println("Task2 is obtained")
   time.Sleep(10*time.Millisecond)
   fmt.Println("Task2 completed")
}

func main(){
   var wg sync.WaitGroup
   wg.Add(1)
   go task1(&wg)
   wg.Add(1)
   go task2(&wg)
   wg.Wait()
}