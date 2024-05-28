package cuckoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	bf, _ := New(WithSize(100, 4))

	testCases := []struct {
		input []byte
		want  bool
	}{
		{input: []byte("hello"), want: true},
		{input: []byte(""), want: true},
	}

	for _, tt := range testCases {
		got := bf.Insert(tt.input)
		assert.Equal(t, got, tt.want, "cuckoo->insert")
	}

	newbf, _ := New(WithSize(10, 4))

	for range 9 {
		newbf.Insert([]byte("hi"))
	}

	assert.Equal(t, newbf.Insert([]byte("hi")), false, "cuckoo->fullbucket insert")
}

func Test_InitialisationError(t *testing.T) {
	_, err := New(WithSize(0, 1))
	assert.NotNil(t, err, "bucket size error")

	_, err = New(WithSize(1, 0))
	assert.NotNil(t, err, "slot size error")

	_, err = New(WithSize(1, 1), WithKicks(0))
	assert.NotNil(t, err, "kick error")
}

func Test_Kicks(t *testing.T) {
	bf, _ := New(WithSize(100, 4), WithKicks(100))
	assert.Equal(t, bf.Kicks(), uint(100))

	nbf, _ := New(WithSize(100, 4))
	assert.Equal(t, nbf.Kicks(), uint(50))
}

func Test_MemberOf(t *testing.T) {
	bf, _ := New(WithSize(100, 4))

	tests := []struct {
		name  string
		input []byte
		want  bool
	}{
		{name: "positive case", input: []byte("hello"), want: true},
		{name: "negative case", input: []byte("bye"), want: false},
		{name: "repeated insertion", input: []byte("repeat"), want: true},
		{name: "empty filter", input: []byte("empty"), want: false},
	}

	bf.Insert(tests[0].input)
	bf.Insert(tests[2].input)
	bf.Insert(tests[2].input)

	for _, tt := range tests {
		got := bf.MemberOf(tt.input)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func Test_Delete(t *testing.T) {
	bf, _ := New(WithSize(100, 4))

	tests := []struct {
		name  string
		input []byte
	}{
		{name: "positive case", input: []byte("hello")},
		{name: "non-existent element", input: []byte("bye")},
		{name: "multiple deletions 1", input: []byte("item1")},
		{name: "multiple deletions 2", input: []byte("item2")},
		{name: "repeated deletion", input: []byte("repeat")},
		{name: "delete from empty filter", input: []byte("empty")},
	}

	bf.Insert(tests[0].input)
	bf.Insert(tests[2].input)
	bf.Insert(tests[3].input)
	bf.Insert(tests[4].input)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name != "non-existent element" && tt.name != "delete from empty filter" {
				bf.Delete(tt.input)
			}

			got := bf.MemberOf(tt.input)
			want := false
			if tt.name == "non-existent element" || tt.name == "delete from empty filter" {
				bf.Delete(tt.input)
				got = bf.MemberOf(tt.input)
			}

			assert.Equalf(t, want, got, tt.name)

			if tt.name == "repeated deletion" {
				bf.Delete(tt.input)
				got = bf.MemberOf(tt.input)
				assert.Equalf(t, want, got, tt.name)
			}
		})
	}
}

func Test_BucketPop(t *testing.T) {
	bf, _ := New(WithSize(100, 4))

	t.Run("correct bucket count after initialization", func(t *testing.T) {
		got := bf.BucketPop()
		assert.Equal(t, uint(100), got, "initial bucket count")
	})

	t.Run("bucket count after insertions", func(t *testing.T) {
		bf.Insert([]byte("item1"))
		bf.Insert([]byte("item2"))
		got := bf.BucketPop()
		assert.Equal(t, uint(100), got, "bucket count after insertions")
	})

	t.Run("bucket count after deletions", func(t *testing.T) {
		bf.Delete([]byte("item1"))
		bf.Delete([]byte("item2"))
		got := bf.BucketPop()
		assert.Equal(t, uint(100), got, "bucket count after deletions")
	})
}
