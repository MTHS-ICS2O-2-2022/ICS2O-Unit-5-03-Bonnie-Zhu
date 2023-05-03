// Copyright (c) 2023 Bonnie Zhu All rights reserved
//
// Created by: Bonnie zhu
// Created on: May 2023
// This program determines age restrictions
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ageInput string
	fmt.Print("Enter your age: ")
	fmt.Scanln(&ageInput)

	age, err := strconv.ParseFloat(ageInput, 64)
	switch {
	case err != nil:
		fmt.Println("Invalid input, please enter a number.")
	case age >= 17:
		fmt.Println("You can see an R rated movie alone.")
	case age >= 13:
		fmt.Println("You can see a PG-13 movie alone.")
	case age >= 6:
		fmt.Println("You can see a G or PG movie alone.")
	default:
		fmt.Println("You're too young to see anything.")
	}
}
