package compress

import (
	"hash/crc32"
	"hash/crc64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CRC32(t *testing.T) {
	tests := []struct {
		input    []byte
		expected uint32
	}{
		{[]byte("hi"), crc32.ChecksumIEEE([]byte("hi"))},
		{[]byte("123921213"), crc32.ChecksumIEEE([]byte("123921213"))},
		{[]byte("hello world"), crc32.ChecksumIEEE([]byte("hello world"))},
		{[]byte(""), crc32.ChecksumIEEE([]byte(""))},
	}

	for _, tt := range tests {
		result := uint32(CRC32(tt.input))
		assert.Equalf(t, result, tt.expected, "")
	}
}

func Test_CRC64(t *testing.T) {
	table := crc64.MakeTable(crc64.ISO)

	tests := []struct {
		input    []byte
		expected uint64
	}{
		{[]byte("hi"), crc64.Checksum([]byte("hi"), table)},
		{[]byte("123921213"), crc64.Checksum([]byte("123921213"), table)},
		{[]byte("hello world"), crc64.Checksum([]byte("hello world"), table)},
		{[]byte(""), crc64.Checksum([]byte(""), table)},
	}

	for _, tt := range tests {
		result := uint64(CRC64(tt.input))
		assert.Equalf(t, result, tt.expected, "")
	}
}
