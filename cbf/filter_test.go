package countingBloom

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	t.Parallel()

	cbf := New(WithSize(1000))

	testCases := []struct {
		input []byte
	}{
		{input: []byte("hello")},
		{input: []byte("")},
		{input: nil},
	}

	for _, tt := range testCases {
		cbf.Insert(tt.input)
		assert.Greaterf(t, len(cbf.freqHashMap), 0, "cbf->insert")
	}
}

func Test_MemberOf(t *testing.T) {
	t.Parallel()

	cbf := New(WithSize(0))

	testCases := []struct {
		key       int
		input     []byte
		expResult bool
	}{
		{1, []byte("hi"), true},
		{2, []byte("bye"), false},
		{3, []byte(uuid.NewString()), true},
		{4, []byte(""), false},
	}

	for _, tt := range testCases {
		if tt.key != 2 && tt.key != 4 {
			cbf.Insert(tt.input)
			result := cbf.MemberOf(tt.input)
			assert.Equalf(t, tt.expResult, result, "cbf->membershipCheck")

			continue
		}

		result := cbf.MemberOf(tt.input)
		assert.Equalf(t, tt.expResult, result, "cbf->membershipCheck")
	}
}

func Test_Delete(t *testing.T) {
	t.Parallel()

	cbf := New(WithSize(10))

	testCases := []struct {
		input      []byte
		isMemberOf bool
	}{
		{[]byte("hi"), false},
		{[]byte(uuid.NewString()), false},
		{[]byte(""), false},
	}

	for _, tt := range testCases {
		cbf.Insert(tt.input)
		cbf.Delete(tt.input)
		assert.Equalf(t, tt.isMemberOf, cbf.MemberOf(tt.input), "cbf->delete")
	}
}

func Test_Flush(t *testing.T) {
	t.Parallel()

	cbf := New(WithSize(10))

	preSeed := len(cbf.freqHashMap)
	cbf.Insert([]byte("hi"))
	postSeed := len(cbf.freqHashMap)
	cbf.Flush()
	postFlush := len(cbf.freqHashMap)

	assert.GreaterOrEqualf(t, postSeed, preSeed, "cbf->flush")
	assert.Less(t, postFlush, postSeed, "cbf->flush")
	assert.Equalf(t, preSeed, postFlush, "cbf->flush")
	assert.Greater(t, cbf.lastVacuumedAt, time.Time{}, "cbf->flush")
}
