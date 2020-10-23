package main
import (
      "fmt"
      "strconv"
     )


type person interface{
   display() string
}


type student struct{
    id int
    name,role string   
}

type teacher struct{
    id int
    name,role string
}

func (s student) display() string{
    var s1 string=strconv.Itoa(s.id)
    return "Displaying "+s.name+" "+s.role+ " "+s1
}

func (t teacher) display() string{
     var s1 string=strconv.Itoa(t.id)
     return "Displaying "+t.name+"  "+t.role+" "+s1
}

func show(p person){
    fmt.Println(p.display())
}

func main(){
   s:=student{1234,"Bapan Biswas","Student"}
   t:=teacher{1233,"Ross Tiler","Teacher"}
   show(t)
   show(s)
   
}