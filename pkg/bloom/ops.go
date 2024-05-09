package bloom

func (f *Filter) InsertAt(v []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()

	segment1, segment2 := murmum3_128(v)
	indexes := f.hashToBitsetIdx([]uint64{segment1, segment2, fnv1a(v)}...)

	for _, idx := range indexes {
		f.bSet.Set(idx)
	}
}

func (f *Filter) MemberOf(v []byte) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	segment1, segment2 := murmum3_128(v)
	indexes := f.hashToBitsetIdx([]uint64{segment1, segment2, fnv1a(v)}...)

	for _, idx := range indexes {
		if f.bSet.Test(idx) {
			return true
		}
	}

	return false
}
