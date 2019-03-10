package main

import (
	"fmt"

	"github.com/techbrosource/go-validators/utils"
)

// User data
type User struct {
	ID   int32  `json:"id" validate:"number"`
	Name string `json:"name" validate:"string,min=2,max=10"`
}

func main() {
	user := User{
		ID:   0,
		Name: "Mr. Steve Jobs",
	}
	response := utils.Validate(user)
	fmt.Println("Errors : ")
	for i := 0; i < len(response.GetFields()); i++ {
		fmt.Print("field '" + response.GetFields()[i].GetName() + "' " + response.GetFields()[i].GetError().Error() + "\n")
	}
}
