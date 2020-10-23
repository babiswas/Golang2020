package main
import "fmt"

type shape interface{
  perimeter() int
  area() int
  getname() string
}

type rectangle struct{
   length,breadth int
   name string
}

type square struct{
   length,breadth int
   name string
}

func (r rectangle) area() int{
    return r.length*r.breadth
}

func (r rectangle) perimeter() int{
   return 2*(r.length+r.breadth)
}

func (r rectangle) getname() string{
   return r.name
}

func (s square) getname() string{
   return s.name
}


func (s square) perimeter() int{
   return 2*(s.length+s.breadth)
}

func (s square) area() int{
    return s.length*s.breadth
}

func main(){
   r1:=rectangle{length:12,breadth:13,name:"House"}
   s1:=square{length:11,breadth:10,name:"Gift"}
   fmt.Println(s1.area())
   fmt.Println(s1.perimeter())
   fmt.Println(r1.area())
   fmt.Println(r1.perimeter())
}