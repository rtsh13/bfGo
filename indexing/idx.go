package indexing

func ModuloBiasing(hashes []uint, n uint) []uint {
	indexes := make([]uint, 0, len(hashes))

	for _, h := range hashes {
		indexes = append(indexes, h%n)
	}

	return indexes
}
