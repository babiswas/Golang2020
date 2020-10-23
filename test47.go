package main
import (
       "fmt"
       "bufio"
       "os"
      )

func main(){
  m:=make(map[string]string,30)
  for i:=0;i<5;i++{
    scanner:=bufio.NewScanner(os.Stdin)
    fmt.Println("Enter key")
    scanner.Scan()
    s1:=scanner.Text()
    fmt.Println("Enter value")
    scanner.Scan()
    s2:=scanner.Text()
    m[s1]=s2  
  }

  for key,value:=range m{
    fmt.Println(key,value)
  }

  delete(m,"hello")
  fmt.Println(m)
}