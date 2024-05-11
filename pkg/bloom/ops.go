package bloom

// inserts the input to the bucket
func (f *Filter) Insert(v []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()

	indexes := f.hashToBitsetIdx(v)

	for _, idx := range indexes {
		// set bit to true for the given idx
		f.bucket.Set(idx)
	}
}

/*
verifies the membership of the input in the bucket.
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

// flushes all the set bits in the bucket
func (f *Filter) Flush() {
	f.bucket = f.bucket.ClearAll()
}

// Count (number of set bits). Also known as "popcount" or "population count".
func (f *Filter) PopCnt() uint {
	return f.bucket.Count()
}
