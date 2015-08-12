package main

import "fmt"

func main() {
	//var arr [5]int = [5]int {1 , 2,3 ,4 ,5}
	m := map[string]string{
		"hello": "world",
		"agile" : "scrum",
	}
	// m["win"] = "golang" //add key
	// delete(m , "win")
	// fmt.Println(m["hello"])
	// fmt.Println(m["win"])

	// fmt.Println(m)

	for k , v := range m{
		// if k == "hello"{
		// 	fmt.Println("Hello Gopher")
		// }
		switch k {
			case k == "agile" : 
				fmt.Println("Hello Scrum")
				fallthrough
			case k == "hello" : 
				fmt.Println("World")
				fallthrough
			default : 
				fmt.Println("Default")
		}
		fmt.Println(k ,v)
	}
}