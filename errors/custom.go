package errors

import "fmt"

type BloomSize struct {
	Size uint
}

func (i BloomSize) Error() string {
	return fmt.Sprintf("invalid bloom filter size : [%v] provided", i.Size)
}

type BloomKicks struct {
	Kick uint
}

func (i BloomKicks) Error() string {
	return fmt.Sprintf("invalid kick size : [%v] provided", i.Kick)
}

type BloomSlots struct {
	Slot uint
}

func (i BloomSlots) Error() string {
	return fmt.Sprintf("invalid slots : [%v] provided", i.Slot)
}
