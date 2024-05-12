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

	/*
		POSITIVE FLOW
	*/
	for i := 0; i < 5; i++ {
		value := uuid.NewString()
		bloomInput, _ := json.Marshal(value)
		fmt.Printf("\nkey : [%v] inserted", value)
		ft.Insert(bloomInput)
		input = append(input, value)
	}

	for i := 0; i < 5; i++ {
		bloomInput, _ := json.Marshal(input[i])
		if !ft.MemberOf(bloomInput) {
			fmt.Printf("\nmembership does not exist for key : [%v]", input[i])
		} else {
			fmt.Printf("\nmembership does exist for key : [%v]", input[i])
		}
	}

	/*
		POSSIBLY NEGATIVE FLOW
	*/
	bloomInput, _ := json.Marshal(uuid.NewString())
	if !ft.MemberOf(bloomInput) {
		fmt.Printf("\nmembership does not exist for key : [%v]", uuid.NewString())
	} else {
		fmt.Printf("\nmembership does exist for key : [%v]", uuid.NewString())
	}
}
