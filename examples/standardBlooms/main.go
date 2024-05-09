package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	bloomFt "github.com/rtsh13/bfGo/pkg/bloom"
)

func main() {
	ft := bloomFt.New(bloomFt.WithSize(5))

	input := make([]string, 0)

	for i := 0; i < 5; i++ {
		value := uuid.NewString()
		bloomInput, _ := json.Marshal(value)
		fmt.Printf("\n key : [%v] inserted \n", value)
		ft.InsertAt(bloomInput)
		input = append(input, value)
	}

	for i := 0; i < 5; i++ {
		bloomInput, _ := json.Marshal(input[i])
		if !ft.MemberOf(bloomInput) {
			fmt.Printf("\nmembership does not exist for key : [%v]\n", input[i])
		} else {
			fmt.Printf("\nmembership does exist for key : [%v]\n", input[i])
		}
	}
}
