package main 

import "fmt"

// class Conatc{
// 	public String name
// }

type Contact struct {
	Name string
	Tel string
}

func say(c Contact) string{
	return fmt.Sprintf("Hello %s" , c.Name)
}
func main() {
	contact := Contact{
		Name : "Kocthaphan",
		Tel : "11111111111",
	}
//contact.Tel = "3222222222222" //change 
	fmt.Println(say(contact))
}