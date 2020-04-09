package main

import (
	"fmt"

	models "github.com/atlit06/webservice/Models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
