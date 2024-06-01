## Counting Bloom Filter

Package `cbf` is a standard representation of a space-efficient probabilistic
Counting Bloom Filter. The factory initialization happens through functional options, which validate the arguments provided by the user. It exposes methods for inserting, verifying membership, deleting elements, and flushing datastore.

CBF's datastore - freqHashMap, stores the hashed inputs, and increments the counter of hashed keys.
in the event of a collision. Every input occupies 3 index positions in the datastore, which are calculated via running 2 different hashing algorithms - <i> Fowler–Noll–Vo hash(Fnv1a)</i> and <i>Murmum3 algorithm</i> and finding index within the boundaries of the set via <i>Modulo Biasing.</i> For concurrent safe operations, we rely on read-write exclusive locks.

### Sequence Flow

[![](https://mermaid.ink/img/pako:eNqNk11vgjAUhv8KOVdzQQNUgxDjxbYsM5lZst1tLKahVZpJa0oxOuN_X21B_M64gNPzvOctB3o2kApCIYZCYUWfGJ5JnLeXQcIdfREmaaqY4M7ru8183X877fbQGfGCSmVzNt5lX3CRMT5LuAXGs046G5s8dHnmSx-fp8elzEuJGmB0BmX-3arVgEppUXAdoT3a2oe9Gze7pSDlXEweGC5MA4YGNym6RQ8_wIgTujr5Asc1xoaR1UlzFzXBPzRH7dqgcjeKt0Vxp9etPQmuEnROLKvXdsviUcznrNAn5bx9Np3YYDBIM8FSOhxW9k2VdamEFa3LDOJMTVJRckVlrJEzxfPisjCVNKd8r3aMXMmyVh847Sr0mavBaaWl4EJOZY4Z0VNi_mACKtPKBGIdEix_Ekj4VutwqcTHmqcQ7_ZzQYpylkFs3tWFckGaEdtnKWFKyLEdQjOLLiww_xSi0eg1xBtYQYzCbifwPOSFUej1u5EfurCG2O_1O5HvB0HUCz2Eeqi7deHXWPjbP51fKeM?type=png)](https://mermaid.live/edit#pako:eNqNk11vgjAUhv8KOVdzQQNUgxDjxbYsM5lZst1tLKahVZpJa0oxOuN_X21B_M64gNPzvOctB3o2kApCIYZCYUWfGJ5JnLeXQcIdfREmaaqY4M7ru8183X877fbQGfGCSmVzNt5lX3CRMT5LuAXGs046G5s8dHnmSx-fp8elzEuJGmB0BmX-3arVgEppUXAdoT3a2oe9Gze7pSDlXEweGC5MA4YGNym6RQ8_wIgTujr5Asc1xoaR1UlzFzXBPzRH7dqgcjeKt0Vxp9etPQmuEnROLKvXdsviUcznrNAn5bx9Np3YYDBIM8FSOhxW9k2VdamEFa3LDOJMTVJRckVlrJEzxfPisjCVNKd8r3aMXMmyVh847Sr0mavBaaWl4EJOZY4Z0VNi_mACKtPKBGIdEix_Ekj4VutwqcTHmqcQ7_ZzQYpylkFs3tWFckGaEdtnKWFKyLEdQjOLLiww_xSi0eg1xBtYQYzCbifwPOSFUej1u5EfurCG2O_1O5HvB0HUCz2Eeqi7deHXWPjbP51fKeM)

### Important considerations:
1. The underlying datastore is a hashmap, which isn't as space efficient as a bitset. This has been designed deliberately as it overcomes the limitations of the standard bloom filter and provides edge by:
    - Dynamic deletion since we track the occurrences of collided hashmap index and increment the counter
    - Reducing false positives over time
    - Avoiding saturation, since the elements can be purged out of CBF
    - Higher accuracy in a dynamic environment(ex: network routing tables)

2. Memory Footprint: freqHashMap is built using `map[uint]int` which is pretty much a memory overhead compared to the vanilla implementation. For instance, consider both filters' datastore keeping 10,000 entries:
    - CBF:
        - Each entry roughly takes up space for a `uint` key (8 bytes on a 64-bit system) and an `int` value (8 bytes), plus some variable overhead.
        - Total size can be approximated as 10,000×(8+8)=160,000 bytes

    - SBF:
        - The bitset will need 1 bit for each possible value.
        - For 10,000 values, the total size is 10,000/8=1250 bytes + minimal overhead.
