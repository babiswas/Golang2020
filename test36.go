package main
import (
	"fmt"
	"time"
	"sync"
)

func main(){

	ch1:=make(chan string)
	ch2:=make(chan string)
	go func(){
		time.Sleep(1*time.Second)
		ch1<-"One"
	}()
    wg.Add(2)
	go func(){
		time.Sleep(2*time.Second)
		ch2<-"Two"
	}()
	
	for{
		select{
		case v:=<-ch1:
			fmt.Println(v)
		case m2:=<-ch2:
			fmt.Println(m2)
		default:
			m++
		}

	}

}