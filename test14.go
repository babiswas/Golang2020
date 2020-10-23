package main
import "fmt"

func main(){
	m:=make(map[string]int)
	m["a"]=2
	m["b"]=3
	fmt.Println(m)
	fmt.Println(len(m))
	m["c"]=5
	fmt.Println(m)
	delete(m,"a")
	fmt.Println(m)
	n:=map[string]int{"key1":1,"key2":2}
	fmt.Println(n)
	x,y:= n["key1"]
	fmt.Println(x)
	fmt.Println(y)
	kl:=[5]int{1,2,3,4,5}
	for k,l:=range kl{
		fmt.Println("key value",k,l)
	}
	for key,val:=range m{
		fmt.Println("Key Value",key,val)
	}
}