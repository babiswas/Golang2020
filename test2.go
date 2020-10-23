package main
import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

type Person struct{
	Name string `json:"name"`
	Age string `json:"age"`
	City string `json:"city"`
}

func main(){
	var person[4]Person
	var person1[4]Person

	for i:=0;i<4;i++{
		scanner:=bufio.NewScanner(os.Stdin)
		fmt.Println("Enter the name of the person")
		scanner.Scan()
		name:=scanner.Text()
		fmt.Println("Enter the age of the Person")
		scanner.Scan()
		age:=scanner.Text()
		fmt.Println("Enter the city of the person")
		scanner.Scan()
		city:=scanner.Text()
		person[i]=Person{Name:name,Age:age,City:city}
	}
    fmt.Println("Displaying the struct to json data")
	for i:=0;i<4;i++{
		data,_:=json.Marshal(person[i])
		fmt.Println(string(data))
	}
	fmt.Println("Converting from json data to struct")
	for i:=0;i<4;i++{
		data,_:=json.Marshal(person[i])
		json.Unmarshal(data,&person1[i])
	}

	fmt.Println("Displaying decoding json in Golang")
	for i:=0;i<4;i++{
		fmt.Println(person1[i])
	}
}