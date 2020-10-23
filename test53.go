package main
import (
   "fmt"
   "io/ioutil"
   "os"
   "log"
   "bufio"
)


func CFile(text,filename string){
   fmt.Println("Creating file")
   file,err:=os.Create(filename)
   if err!=nil{
     log.Fatalf("Couldn't create the file")
   }
   data,err:=file.WriteString(text)
   if err!=nil{
      log.Fatalf("Unable to write on file")
   }
   fmt.Println(data)
}


func RFile(filename string){
   fmt.Println("Reading a file")
   file,err:=ioutil.ReadFile(filename)
   if err!=nil{
      log.Fatalf("Unable to read the file")
   }
   str:=string(file)
   fmt.Println(str)
}


func main(){
   var filename string
   fmt.Scanln(&filename)
   input:=bufio.NewReader(os.Stdin)
   outpt,err:=input.ReadString('\n')
   if err!=nil{
      log.Fatalf("Could not read the string")
   }
   CFile(outpt,filename)
   RFile(filename)
}