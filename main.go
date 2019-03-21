package main

import (
	"fmt"

	"github.com/techbrosource/go-validators/utils"
)

// User data
type User struct {
	ID   int32  `json:"id" required:"true"`
	Name string `json:"name" required:"true" minlen:"2" maxlen:"10"`
}

func main() {
	user := User{
		ID:   0,
		Name: "Mr. Steve Jobs",
	}
	response := utils.Validate(user)
	fmt.Println("Errors : ")
	for i := 0; i < len(response); i++ {
		fmt.Print("field '" + response[i].GetName() + "' " + response[i].GetError().Error() + "\n")
	}
}
