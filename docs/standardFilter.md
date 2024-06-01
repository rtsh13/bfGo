## Bloom Filter

Package `bloomFilter` is a standard representation of a space-efficient probabilistic
data structure.  It exposes methods for inserting and verifying membership in an
input within the set. It is very common for an input to occupy two places in the set, which is also referred to as <i>collisions</i>. Therefore, standard bloom filters do not support the deletion of inserted inputs.

The factory initialization happens through functional options, which validate the arguments provided by the user.
[Bitset](https://github.com/bits-and-blooms/bitset) is the underlying datastore for the bloom filter, being the source of truth for data. Each operation is wrapped around <i>read-write exclusive locks</i> for a concurrent environment.
In its current state, every input occupies 3 bits in a set, which is calculated via running 2 different hashing algorithms: <i> Fowler-Noll–Vo hash(Fnv1a)</i> and <i>Murmum3 algorithm</i> and finding index within the boundaries of the set via modulo biasing.</i>

### Sequence Flow

[![](https://mermaid.ink/img/pako:eNqNkl1PwjAUhv_Kcq7EDNxayNguvDDGSCIx0cQLnSF1LayRtUu3kSHZf7e0wPgQ4i66nvd5e9rTnhUkkjKIoChJye45mSmSdRcoFo7-KFcsKbkUztOLVT6uP51u99YZiYKp8mbMsi-mipTnkzem-JQnZG2PhXVfdplEj0TrYrZdYY6xFZ2VFfc3fhALn5zK40pllcItMD6DUv-q7rRg47QInUd4hxr7s6PJZreUtJrLyR0nhSnAUHSR4kt0_wJGgrL66AYO15g0nNZHxf3pQf_wHJRrJ5vsxvGcF1c67uwIOkvwKbFsGxuoH26tgwsZUxnhVPegKTaGMmUZiyHSU0rUdwyxaLSPVKV8XYoEolJVzAUlq1kK0ZTMCx1VOW0beKcyykupxrbFTae7kBPxLmXr0TFEK6ghwkG_hzwPe0EYeMN-6AcuLCHyB8Ne6PsIhYPAw3iA-40LPyaF3_wCYvr2dQ?type=png)](https://mermaid.live/edit#pako:eNqNkl1PwjAUhv_Kcq7EDNxayNguvDDGSCIx0cQLnSF1LayRtUu3kSHZf7e0wPgQ4i66nvd5e9rTnhUkkjKIoChJye45mSmSdRcoFo7-KFcsKbkUztOLVT6uP51u99YZiYKp8mbMsi-mipTnkzem-JQnZG2PhXVfdplEj0TrYrZdYY6xFZ2VFfc3fhALn5zK40pllcItMD6DUv-q7rRg47QInUd4hxr7s6PJZreUtJrLyR0nhSnAUHSR4kt0_wJGgrL66AYO15g0nNZHxf3pQf_wHJRrJ5vsxvGcF1c67uwIOkvwKbFsGxuoH26tgwsZUxnhVPegKTaGMmUZiyHSU0rUdwyxaLSPVKV8XYoEolJVzAUlq1kK0ZTMCx1VOW0beKcyykupxrbFTae7kBPxLmXr0TFEK6ghwkG_hzwPe0EYeMN-6AcuLCHyB8Ne6PsIhYPAw3iA-40LPyaF3_wCYvr2dQ)

### Important considerations:
1. The maximum size of the bitset is
    - (2<sup>32</sup> - 1) : 4294967295 bits for x32 bit architecture
    - (2<sup>64</sup> - 1) : 18446744073709551615 bits for x64 bit architecture

2. While the underlying datastore provides extending of size during <b>saturation</b>, the implementation aims to defer since the datastore operations are agnostic of the input and its hash values.

3. If the datastore capacity is extended, the membership verification of existing inserted elements will fail, as the indexing has a direct correlation with the originally provided size. While there are ways to tackle this scenario, it would, however, be futile considering that
    - would require rehashing all of the existing elements, defeating the purpose of using a Bloom filter for its constant-time insertion and lookup characteristics.
    - computationally expensive and warrants for huge memory footprint