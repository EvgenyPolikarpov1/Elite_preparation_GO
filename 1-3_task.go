package main

import "fmt"

func main() {
	//Task 1
	arr := []string{"Monday" , "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	slice := arr[1:]
	fmt.Println(slice)
	fmt.Println(len(arr),cap(arr))
	fmt.Println(len(slice),cap(slice))
	
	//Task2
	fmt.Println(arr)
	fmt.Println(slice[0])
	fmt.Println(arr[0], slice[1:])

	//Task3
	slice = append(slice, "Friday", "Friday", "Friday", "Friday", "Friday", "Friday", "Friday")
	fmt.Println(slice)


	

}