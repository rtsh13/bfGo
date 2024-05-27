package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	cuckoo "github.com/rtsh13/bfGo/cuckooFilter"
)

func main() {
	cf, err := cuckoo.New(cuckoo.WithSize(10, 4), cuckoo.WithKicks(5))
	if err != nil {
		log.Fatalf("error : [%v] in initializing filter", err.Error())
		return
	}

	input := make([]string, 0)

	///////////////////////////////////////////
	// 				INSERT
	///////////////////////////////////////////
	fmt.Println()
	for i := 0; i < 5; i++ {
		value := uuid.NewString()
		bloomInput, _ := json.Marshal(value)
		fmt.Printf("\nkey : [%v] inserted", value)
		if !cf.Insert(bloomInput) {
			fmt.Printf("\n at full capacity to insert : [%v]", bloomInput)
		}

		input = append(input, value)
	}

	///////////////////////////////////////////
	// 				MEMBERSHIP
	///////////////////////////////////////////
	fmt.Println()
	for i := 0; i < 5; i++ {
		bloomInput, _ := json.Marshal(input[i])
		fmt.Printf("\nmembership : [%v] for key : [%v]", cf.MemberOf(bloomInput), input[i])
	}

	///////////////////////////////////////////
	// 				DELETE
	///////////////////////////////////////////
	fmt.Println()
	for i := 0; i < 5; i++ {
		bloomInput, _ := json.Marshal(input[i])
		fmt.Printf("\nkey : [%v] deleted", input[i])
		cf.Delete(bloomInput)
	}

	///////////////////////////////////////////
	// 				MEMBERSHIP
	///////////////////////////////////////////
	fmt.Println()
	for i := 0; i < 5; i++ {
		bloomInput, _ := json.Marshal(input[i])
		fmt.Printf("\nmembership : [%v] for key : [%v]", cf.MemberOf(bloomInput), input[i])
	}
}
