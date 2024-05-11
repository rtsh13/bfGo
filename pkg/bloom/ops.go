package bloom

func (f *Filter) Flush(v []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()

	indexes := f.hashToBitsetIdx(v)

	for _, idx := range indexes {
		// set bit to true for the given idx
		f.bucket.Set(idx)
	}
}

/*
Verifies the membership of the input in the bitset.
To avoid false positives, since the likelyhood of the
collision is unknown due to variable space, bits for
all the calculated indexes are verified.
*/
func (f *Filter) MemberOf(v []byte) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	indexes := f.hashToBitsetIdx(v)

	for _, idx := range indexes {
		if !f.bucket.Test(idx) {
			return false
		}
	}

	return true
}
