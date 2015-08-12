package contact 

import "fmt"

// class Conatc{
// 	public String name
// }
func main() {
	c := &contact.Contact{
		Name: "Kotchaphan",
	}
	fmt.Println(c)
}
type Contact struct {
	Name string `json:"name"`//`xml:"name"`
	Tel string 	`json:"tel"`	
	Email string `json:"email"`
}

func (c *Contact) String() string{
	return fmt.Sprintf("Hello %s" , c.Name)
}

func Say(c Contact) string{
	return fmt.Sprintf("Hello %s" , c.Name)
}
