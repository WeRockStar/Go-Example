package main

import "fmt"

func main() {	
	var arr [5]int = [5]int {1 , 2,3 ,4 ,5}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// i := 0
	// for i < len(arr){
	// 	fmt.Println(arr[i])
	// 	i++
	// }

	// for i ,v := range arr {
	// 	fmt.Println(i , v)
	// }

	for _,v := range arr {
		fmt.Println(v)
	}
}