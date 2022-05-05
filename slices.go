package main

import "fmt"

func slicetest() {
	fmt.Println("Slice Test")

	s := []string{"a","b"}
	fmt.Println(s, len(s))

	numbers := [5]int{1, 2, 3, 4, 5}

	//Both start and end [end shiould be grater]
	num1 := numbers[2:5]
	fmt.Println( "Len => ", len(s))
	fmt.Println( num1)

	//Only start&end
	//<Start_INDEX> : <END_INDX> -> END_INDEX will ignore the last index
	fmt.Println(numbers[2:], numbers[:len(numbers)], numbers[2: 5])

	//Copy
	src := []int{1, 2, 3, 4, 5}
    dst := make([]int, 5)

	fmt.Println("before Copy ", dst, src)
	copy(dst, src)
	fmt.Println("after Copy ", dst, src)


}