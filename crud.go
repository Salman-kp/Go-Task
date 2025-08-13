package main

import (
	"fmt"
)

func crud() {
	for {
		fmt.Println("\n === CRUD App Meanu ===")
		fmt.Println("1.create")
		fmt.Println("2.read")
		fmt.Println("3.update")
		fmt.Println("4.delete")
		fmt.Println("5.exit")

		var choice int
		fmt.Print("Enter the number : ")
		fmt.Scan(&choice)

		switch choice{
		case 1:
			fmt.Println("Create operation selected")
		case 2:
			fmt.Println("Read operation selected")
		case 3:
			println("Update operation selected")
		case 4:
			println("Deleted operation selected")
		case 5:
			println("exite from this")
		return
		default:
			fmt.Println("enter the number between the 1-5")
		}

	}
}
