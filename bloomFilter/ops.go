package bloom

// inserts the input to the buckets
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
verifies the membership of the input in the buckets.
To avoid false +ve, since the likelihood of the collision is
unknown due to variable space, we defer to total membership check.
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

// flushes all the set bits in the buckets
func (f *Filter) Flush() {
	f.bucket = f.bucket.ClearAll()
}

// Count (number of set bits). Also known as "popcount" or "population count".
func (f *Filter) PopCnt() uint {
	return f.bucket.Count()
}
