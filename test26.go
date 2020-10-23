package main
import (
	"io"
	"net"
	"time"
	"fmt"
)

func handleconn(c net.Conn){
   defer c.Close()
   for{
	   _,err:=io.WriteString(c,"response from server\n")
	   if err!=nil{
		   return
	   }
	   time.Sleep(time.Second) 
	}
}

func main(){

	listener,err:=net.Listen("tcp","localhost:8000")
	if err!=nil{
		fmt.Println("Error Occured")
	}

	for{
		fmt.Println("Server got a request")
		conn,err:=listener.Accept()
		if err!=nil{
			continue
		}
		fmt.Println("Server Accepted Client")
		go handleconn(conn)
	}

}

