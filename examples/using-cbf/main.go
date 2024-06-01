package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	cbf "github.com/rtsh13/bfGo/cbf"
)

func main() {
	ft, _ := cbf.New(cbf.WithSize(100000))

	input := make([]string, 0)

	///////////////////////////////////////////
	// 				INSERT
	///////////////////////////////////////////
	fmt.Println()
	for i := 0; i < 5; i++ {
		value := uuid.NewString()
		bloomInput, _ := json.Marshal(value)
		fmt.Printf("\nkey : [%v] inserted", value)
		ft.Insert(bloomInput)
		input = append(input, value)
	}

	///////////////////////////////////////////
	// 				MEMBERSHIP
	///////////////////////////////////////////
	fmt.Println()
	for i := 0; i < 5; i++ {
		bloomInput, _ := json.Marshal(input[i])
		fmt.Printf("\nmembership : [%v] for key : [%v]", ft.MemberOf(bloomInput), input[i])
	}

	///////////////////////////////////////////
	// 				DELETE
	///////////////////////////////////////////
	fmt.Println()
	for i := 0; i < 5; i++ {
		bloomInput, _ := json.Marshal(input[i])
		fmt.Printf("\nkey : [%v] deleted", input[i])
		ft.Delete(bloomInput)
		fmt.Printf("\nmembership : [%v] for key : [%v]", ft.MemberOf(bloomInput), input[i])
	}
}
