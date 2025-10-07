package main

import "fmt"

func studentmark() {


	var upmark int
	var student string

	std := make(map[string]int)
	std["shamil"]=99
	std["salman"]=99
	std["sinan"]=99

	fmt.Println("Enter the Name : ")
	fmt.Scan(&student)
	
	fmt.Println("Enter the updated mark : ")
	fmt.Scan(&upmark)

	for name:=range std{
		if name==student{
			std[name]=upmark
		}
	}
	fmt.Println("The Details of students : ",std)
}