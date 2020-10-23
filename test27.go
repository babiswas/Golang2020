package main
import (
	"fmt"
	"io"
	"net"
	"os"
)

func main(){
	conn,err:=net.Dial("tcp","localhost:8000")
	defer conn.Close()
	if err!=nil{
		fmt.Println("Error Occured")
	}
	mustcopy(os.Stdout,conn)
}


func mustcopy(dst io.Writer,src io.Reader){
   _,err:=io.Copy(dst,src)
   fmt.Println(err)
   if err!=nil{
	   fmt.Println("Fatal Error")
   }
}