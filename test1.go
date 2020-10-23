package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"
)

func get_data()string{
	var data string
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data=scanner.Text()
	return data
}



type Person struct{
	Name string  `json:"name"`
	Age string   `json:"age"`
	Degree string `json:"degree"`
}

func main(){
	var person [3]Person
	for i:=0;i<3;i++{
		fmt.Println("Enter the name of the person")
		name:=get_data()
		fmt.Println("Enter the age of the person")
		age:=get_data()
		fmt.Println("Enter the degree of the person")
		degree:=get_data()
		person[i]=Person{Name:name,Age:age,Degree:degree}
	}

	for i:=0;i<3;i++{
		data,err:=json.Marshal(person[i])
		if err!=nil{
			fmt.Println("Error Occured")
		}
		fmt.Println(string(data))
	}

}
