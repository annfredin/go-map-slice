package main

import (
	"fmt"
)



func maptest() {
	fmt.Println("\n Map Test")

	//ERR: assignment to entry in nil map
	// var emp map[string]string //not ok
	 emp := map[string]string{} //ok
	//  emp := make(map[string]string) // ok
    emp["test"] = "Test"
	fmt.Println(emp)


	//setting key
	emp["name"] = "Fredin"
	//retriving key
	getOp := emp["desc"]
	fmt.Println("getOp: ", getOp)

	//exist check
	getOp, ok := emp["desc"]
	if !ok {
		fmt.Println("KEY-> desc not found")
	}

	fmt.Println(emp)
	fmt.Println("Before Deleting test ", emp)
	delete(emp, "test")
	fmt.Println("After Deleting test ", emp)

	//length
	fmt.Println("Length ", len(emp))


	//REFERENCE TYPE CHK
	  //Declare
	  e1 := make(map[string]int)
	  e1["sal"] = 2000
  
	  e2 := e1
	  e2["sal"] = 5000
	fmt.Println("e1: ", e1, " e2: ", e2)

	

}