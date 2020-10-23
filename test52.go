package main
import (
        "fmt"
        "io/ioutil"
        "os"
        "log"
)

func fileWrite(){
   file,err:=os.Create("file2.txt")
   defer file.Close()
   if err!=nil{
      log.Fatalf("Couldn't create the file")
   }
   data,err:=file.WriteString("Hello world How are you")
   if err!=nil{
     log.Fatalf("Couldn't wtrite in the file")
   }
   fmt.Println(data)
}

func readfile(){
    file,err:=ioutil.ReadFile("file2.txt")
    if err!=nil{
       log.Fatalf("File reading failed")
    }
    str:=string(file)
    fmt.Println(str)
}

func main(){
   fileWrite()
   readfile()
}