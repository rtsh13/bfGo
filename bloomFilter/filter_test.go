package bloom

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	t.Parallel()

	bloomFt, _ := New(WithSize(100))

	preSetPopCnt := bloomFt.PopCnt()
	bloomFt.Insert([]byte("hi"))
	postSetPopCnt := bloomFt.PopCnt()

	assert.GreaterOrEqualf(t, postSetPopCnt, preSetPopCnt, "filter->insert")
}

func Test_MemberOf(t *testing.T) {
	t.Parallel()

	bloomFt, _ := New(WithSize(100))

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
			bloomFt.Insert(tt.input)
			result := bloomFt.MemberOf(tt.input)
			assert.Equalf(t, tt.expResult, result, "filter->memberOf")

			continue
		}

		result := bloomFt.MemberOf(tt.input)
		assert.Equalf(t, tt.expResult, result, "filter->memberOf")
	}
}

func Test_Flush(t *testing.T) {
	t.Parallel()

	bloomFt, _ := New(WithSize(100))

	preSeed := bloomFt.PopCnt()
	bloomFt.Insert([]byte("hi"))
	postSeed := bloomFt.PopCnt()
	bloomFt.Flush()
	postFlush := bloomFt.PopCnt()

	assert.GreaterOrEqualf(t, postSeed, preSeed, "filter->flush")
	assert.Less(t, postFlush, postSeed, "filter->flush")
	assert.Equalf(t, preSeed, postFlush, "filter->flush")
}
