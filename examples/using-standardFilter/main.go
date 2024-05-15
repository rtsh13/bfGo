package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	bloomFt "github.com/rtsh13/bfGo/bloomFilter"
)

func main() {
	ft := bloomFt.New(bloomFt.WithSize(100))

	input := make([]string, 0)

	//////////////////////////
	//		  insertion in BF
	/////////////////////////
	fmt.Println()
	for i := 0; i < 5; i++ {
		value := uuid.NewString()
		bloomInput, _ := json.Marshal(value)
		fmt.Printf("\nkey : [%v] inserted", value)
		ft.Insert(bloomInput)
		input = append(input, value)
	}

	//////////////////////////
	//		 membership check
	/////////////////////////
	fmt.Println()
	for i := 0; i < 5; i++ {
		bloomInput, _ := json.Marshal(input[i])
		fmt.Printf("\nmembership : [%v] for key : [%v]", ft.MemberOf(bloomInput), input[i])
	}

	//////////////////////////
	//	   new key membership
	/////////////////////////
	fmt.Println()
	key := uuid.NewString()
	bloomInput, _ := json.Marshal(key)
	fmt.Printf("\nmembership : [%v] for key : [%v]", ft.MemberOf(bloomInput), key)
}
