## Cuckoo Bloom Filter

Package `cuckooFilter` is a representation of a space-efficient probabilistic
Cuckoo Bloom Filter. The factory initialization happens through functional options, which validate the arguments provided by the user. It exposes methods for inserting, verifying membership, deleting elements, rendering configured kicks, and count of hashtables.

Cuckoo's datastore is a 2D representation of elements being persisted, where we create `m` hashtables and `n` possible slots for each hashtable. The operations of the hashtable are decoupled from those of the filter. `k` is used to define the possible iterations of open addressing that shall be executed in the event of a full collision.

The idea behind Cuckoo Filter's datastore is aided via [Cuckoo Hashing](https://en.wikipedia.org/wiki/Cuckoo_hashing). We create one or more hashtables, each having a predefined capacity to store the fingerprints (fp). Consider fp to be a singular compressed value that identifies an element uniquely. During insertion, we find hashtables to store the element in one of the slots. The calculated fp, vis-à-vis the Cyclic Redundancy Check (CRC) algorithm, is stored in the desired position if it is empty. However, in the event of a full collision, where all the hashtables' slots are occupied, we execute [open addressing or closed hashing](https://en.wikipedia.org/wiki/Open_addressing) to find alternate positions.

Open Address calculates the next hashtable and slot where the element needs to be inserted by evicting the ones that are occupying the position and running recursively. However, this is contained by the defined `kicks(k)` to avoid an infinite loop in case of total saturation of the filter.

Every input occupies 2 hashtable's slots in the datastore, which is calculated by running <i> Fowler-Noll-Vo hash(Fnv1a)</i> and its symmetric hash via XOR operation, followed by Modulo Biasing. For concurrent safe operations, we rely on read-write exclusive locks.

### 1. Insertion

Below is the sequential flow for insertion on an input in CF without diving into open addressing.

[![](https://mermaid.ink/img/pako:eNqNVGtr2zAU_StCn7LhGD-SODFskGYUCusGadmH1R0othKL2pKR5KxuyH-fHk5sQdlmMEjnnnN0dX19TzBnBYYpFBJJ_IWgA0f19BhlFKinIBznkjAKvm4t8vTxGUynn8EdFZjLjFrUiHsMnCymH0KbVhr-vvmBqhTkrG44FsLfbDeL2cTEPwx8w-r5N53EI8Ej08Cac9RNDM2R6ZjRlWGqXvAJlEiU_i09hmhiwyO-IhhupLiR5obg198FkREQWuBXLNLLQknNitCDf8-KtmI3BAm1mzyVoadUzx7Y-7sNa6m-5bguVq9NNyXOXx4qJsWYYSs6xMZVdb6DxFwx7y6Z9XsgS87aQ3k5yRW7IjeL9RGRCu1IRWSXWhiQPRAqBIgAyIYr7Dq-KzfG3_O8bTodUsn1PlcTDzATtvBvIkvbAq75YGEcH3mrGmOLZcspkGrzv6l8w6_STYS6yeTmtlTRTNj1vajfKfs_q3uLKoHNsZQZa1VIjsdHc3udvSYOdudxS4yaQXuyBtN1UeifQ3XcQNPlMQTVIgNqMnDhq3n_3_ZBp7SAcZsT2CGBC6AmATFsPRPU0W2lqgQ9WGNeI1KoOWL6NIOyxDXOYKqWBeIvGczoWfFQK9lDR3OYancPmiaFqTnCg21TDEPoiuKCSMbv7Zgy08qDDaI_GasvNmoL0xN8hek0Tvw4CZJ4tgzD5WIeLDzYKXgR-7MkWgZBEM7j1XIezM8efDMWoR-G0WqZLFbzONGy6PwHNJeQpQ?type=png)](https://mermaid.live/edit#pako:eNqNVGtr2zAU_StCn7LhGD-SODFskGYUCusGadmH1R0othKL2pKR5KxuyH-fHk5sQdlmMEjnnnN0dX19TzBnBYYpFBJJ_IWgA0f19BhlFKinIBznkjAKvm4t8vTxGUynn8EdFZjLjFrUiHsMnCymH0KbVhr-vvmBqhTkrG44FsLfbDeL2cTEPwx8w-r5N53EI8Ej08Cac9RNDM2R6ZjRlWGqXvAJlEiU_i09hmhiwyO-IhhupLiR5obg198FkREQWuBXLNLLQknNitCDf8-KtmI3BAm1mzyVoadUzx7Y-7sNa6m-5bguVq9NNyXOXx4qJsWYYSs6xMZVdb6DxFwx7y6Z9XsgS87aQ3k5yRW7IjeL9RGRCu1IRWSXWhiQPRAqBIgAyIYr7Dq-KzfG3_O8bTodUsn1PlcTDzATtvBvIkvbAq75YGEcH3mrGmOLZcspkGrzv6l8w6_STYS6yeTmtlTRTNj1vajfKfs_q3uLKoHNsZQZa1VIjsdHc3udvSYOdudxS4yaQXuyBtN1UeifQ3XcQNPlMQTVIgNqMnDhq3n_3_ZBp7SAcZsT2CGBC6AmATFsPRPU0W2lqgQ9WGNeI1KoOWL6NIOyxDXOYKqWBeIvGczoWfFQK9lDR3OYancPmiaFqTnCg21TDEPoiuKCSMbv7Zgy08qDDaI_GasvNmoL0xN8hek0Tvw4CZJ4tgzD5WIeLDzYKXgR-7MkWgZBEM7j1XIezM8efDMWoR-G0WqZLFbzONGy6PwHNJeQpQ)

<br></br>

### 2. Open Addressing
Below is the sequential flow for insertion on an input in CF by linear probing until max (kicks).

[![](https://mermaid.ink/img/pako:eNp9kk2P2jAQhv_KyMcqIAIkWXyotGpVCa22h1L10E0PbjwhFo4d2Q4si_jvdexlCVu1Pnk-3scz4zmRSnMklFjHHH4WbGtYO9nPSwX-PH34BZPJR9AdqnvODVor1LZUMRok72JwirGxeu3Q-MwHUe3sRTucsT8kblBi5e4V3xxYR19NMExx3YJQHJ_B38FK7RKwPgf2TPZor8gbQmB-w4bZxnvWg56-2lB3P5hMoGKy6uXQhsJDfGKMuxUH3qcGq91XPGx8EZZGM4iHqizU2gDbMyHZbyGFO477vZEG2HfTI4V1HcQXncQEDLreKHA-_j_9eIKBo_RfqEorJ1SPsIvz_-f0vzBpYzVMypgNtce8VVMPCRFwvvQ1dBDU_rOjJ2CuLpKQFk3LBPdLFpajJK7BFktC_ZUzsytJqc4-j_VOb46qInToOyFG99uG0PBsQvqOXzf0zYtcOG0e4w6HVU5Ix9RPrdsLxpuEnsgzoZMsneb5YrHIVvPZLL3L0oQcvXuVT4vZMs3zNFvdzRer4pyQl0BIp1laLGdFnmWFDy7ny_Mfy0QEIA?type=png)](https://mermaid.live/edit#pako:eNp9kk2P2jAQhv_KyMcqIAIkWXyotGpVCa22h1L10E0PbjwhFo4d2Q4si_jvdexlCVu1Pnk-3scz4zmRSnMklFjHHH4WbGtYO9nPSwX-PH34BZPJR9AdqnvODVor1LZUMRok72JwirGxeu3Q-MwHUe3sRTucsT8kblBi5e4V3xxYR19NMExx3YJQHJ_B38FK7RKwPgf2TPZor8gbQmB-w4bZxnvWg56-2lB3P5hMoGKy6uXQhsJDfGKMuxUH3qcGq91XPGx8EZZGM4iHqizU2gDbMyHZbyGFO477vZEG2HfTI4V1HcQXncQEDLreKHA-_j_9eIKBo_RfqEorJ1SPsIvz_-f0vzBpYzVMypgNtce8VVMPCRFwvvQ1dBDU_rOjJ2CuLpKQFk3LBPdLFpajJK7BFktC_ZUzsytJqc4-j_VOb46qInToOyFG99uG0PBsQvqOXzf0zYtcOG0e4w6HVU5Ix9RPrdsLxpuEnsgzoZMsneb5YrHIVvPZLL3L0oQcvXuVT4vZMs3zNFvdzRer4pyQl0BIp1laLGdFnmWFDy7ny_Mfy0QEIA)