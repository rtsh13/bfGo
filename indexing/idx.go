package indexing

/*
Once the hashes are generated by hashing module, we need to find one/many
index position(s) that will flush/seek the data. This is deliberately done
to minimize the false postives resulted by collisions.

	index = {hash(i) % bitset_size}

For each hash h, it calculates the modulo operation to ensure that the index
falls within the range of the bitset size
*/
func ModuloBiasing(hashes []uint, n uint) []uint {
	indexes := make([]uint, 0, len(hashes))

	for _, h := range hashes {
		indexes = append(indexes, h%n)
	}

	return indexes
}
