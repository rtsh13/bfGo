# bfGo 

[![codecov](https://codecov.io/gh/rtsh13/bfGo/graph/badge.svg?token=U17GKADZU6)](https://codecov.io/gh/rtsh13/bfGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/rtsh13/bfGo)](https://goreportcard.com/report/github.com/rtsh13/bfGo)
[![Maintainability](https://api.codeclimate.com/v1/badges/f603603bd14764748fe3/maintainability)](https://codeclimate.com/github/rtsh13/bfGo/maintainability)
<a href="https://opensource.org/licenses/Apache-2.0"><img src="https://img.shields.io/badge/License-Apache_2.0-blue.svg"></a>

`bfGo` is a library written in Golang to provide end user/application to spin up bloom filters from scratch. It provides configurable parameters and multiple variants such as:

1. Standard Bloom Filter
    - SBF relies on underlying bitset as a datastore and provisions Insertion and Membership verification of inputs
2. Counting Bloom Filter
    - CountingBF relies on underlying frequencyHashmap as a datastore and provisions Insertion, Membership verification and deletion of inputs
3. Cuckoo Bloom Filter
    - CuckooBF relies on underlying hashtables and slots as a datastore and provisions Insertion, Membership verification, Deletion and other auxilary operations

In addition, bfGo has `Hashing`, `Indexing` and `Compression` modules that aid the lifecyle of an input in bloom filters. bfGo uses consistent hashing algorithms are deterministic such as <i>Fowler-Noll–Vo hash(Fnv1a)</i> and <i>Murmum3 algorithm</i>, whereas <i>Module Biasing</i> is used for searching indexes and <i>Cyclic Redundancy Checksum</i> algorithm is used for compression, primarly for Cucoko filters.


Each of the filters have their own designed datastores which are compact and memory efficient, keeping footprint as minimal as possible. For more details on how these are designed, please redirect to [/docs](https://github.com/rtsh13/bfGo/tree/development/docs) for a better understanding.


## Usage Tip

1. Use standard bloom filters when you need a space-efficient probabilistic data structure to test membership of an element in a set
    - Target usecases: spell checkers, web caches, or distributed systems.
    - Parameters:
       - Insertion time and Membership verification is directly proportional to input size. 
        - Operation on the datastore runs on constant time
        - If the expected inputs are huge in size, it would be recommended to use strong algorithms that are blazing fast to find hashes.

2. Use counting bloom filters when you need to support both insertion and deletion of elements while maintaining a compact data structure. 
    - Target usecases: network traffic monitoring, cache management, or stream processing.
    - Parameters: 
        - Insertion time and Membership verification is directly proportional to input size. 
        - Operation on the datastore runs on constant time
        - If the expected inputs are huge in size, it would be recommended to use strong algorithms that are blazing fast to find hashes.

3. Use cuckoo bloom filters when you need to handle high insertions rates and want to minimize false positives. 
    - Target usecases: such as network security, malware detection, or DNA sequence analysis.
    - Parameters: 
        - Slot size has minimal impact in performance, but allocations and total insertion time is directly proportional to input size.
        - Operation on the datastore runs on constant time
        - Cuckoo filters generally work best when you have filter size is 10x or greater than the input size and count.
        - If the expected inputs are huge in size, it would be recommended to use strong algorithms that are blazing fast to find hashes.

## Performance
 
Each of the filters have been benchmarked and tested throughly to understand the performance and correlation of input size, filter size and hashing functions. The factors used to conduct the tests are :[INPUT SIZE/FILTER SIZE/SLOT SIZE] in the exact order.

Based on below benchmarks, it can be concluded that with ideal balance between filter, slot and input size, where filter size is always 10x or greater than the input size(which is not too large ie; <1000 bytes), would result in operations being executed in approximately:

1. `6μs`(insertion) and  `1μs`(membership verification) in SBF
2. `7μs`(insertion) and `1μs`(membership verification) in Counting BF
3. `6μs`(insertion) and `0.5845μs`(membership verification) in Cuckoo BF

| Benchmark                                     | Operations | Time/op      | Memory/op | Allocations/op |
|------------------------------------------     |------------|-----------   |-----------|----------------|
| Standard BF Insert [1000/1000]                | 168,452    | 6,729 ns     | 1,144 B   | 3              |
| Standard BF Membership [1000/1000]            | 926,133    | 1,255 ns     | 120 B     | 2              |
| CBF Insert [1000/1000]                        | 155,170    | 7,119 ns     | 1,168 B   | 4              |
| CBF Membership [1000/1000]                    | 900,470    | 1,339 ns     | 144 B     | 3              |
| Cuckoo Filter Insert [1000/1000/10]           | 156175     | 13091 ns/op  | 1834 B/op | 51             |
| Cuckoo Filter Insert [1000/100000/10]         | 180880     | 6317 ns/op   | 1040 B/op | 2              |
| Cuckoo Filter Membership [1000/1000/10]       | 2009852    | 584.1 ns/op  | 16 B/op   | 1              |
| Cuckoo Filter Membership [1000/100000/10]     | 2074994    | 584.5 ns/op  | 16 B/op   | 1              |
| Cuckoo Filter Membership [1000/1000/1000]     | 2082231    | 582.0 ns/op  | 16 B/op   | 1              |
| Cuckoo Filter Membership [1000/1000/10000]    | 2018246    | 570.1 ns/op  | 16 B/op   | 1              |


## Contributing 
Contributions and collaborations are much welcomed! Please head over to the [guidelines](https://github.com/rtsh13/bfGo/blob/development/CONTRIBUTING.md) and raise a PR for the change.