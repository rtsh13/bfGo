package compress

import (
	"encoding/binary"
	"hash/crc32"
	"hash/crc64"
	"unsafe"
)

const (
	arch32ByteSize = 4
	arch64ByteSize = 8
)

func CRC32(input []byte) uint {
	h := crc32.NewIEEE()
	h.Write(input)

	return uint(h.Sum32())
}

func CRC64(input []byte) uint {
	table := crc64.MakeTable(crc64.ISO)
	h := crc64.New(table)
	h.Write(input)

	return uint(h.Sum64())
}

func ToByteArray(input uint) []byte {
	switch unsafe.Sizeof(input) {
	case arch32ByteSize:
		buff := make([]byte, arch32ByteSize)
		binary.BigEndian.PutUint32(buff, uint32(input))

		return buff
	default:
		buff := make([]byte, arch64ByteSize)
		binary.BigEndian.PutUint64(buff, uint64(input))

		return buff
	}
}
