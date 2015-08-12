package main
import "fmt"
//var s = "hello golang"
func main() {
	//var s = "hello golang" check type 
	//s = 1
	//s:=1
	//s := strconv.Itoa(1)
	//s := []byte("hello golang")
	// := dynamic programming no static type
	//var s = "hello golang"
	//s := []byte("hello golang")
	//var b bool = true //false
	//var f float32 = 1.232423
	var s  string = "kotchaphan musangn"
	var i [5]int = [5]int {1 , 3 , 2 ,2} //array
	var j = i[0:2] //substring
	var str = s[0:5] //substring 
	fmt.Println(j)
	fmt.Println(len(i))
	fmt.Println(str)
}