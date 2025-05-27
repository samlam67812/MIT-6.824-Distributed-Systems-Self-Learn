### Why build distributed systems
- increase capacity via parallel process
- tolerate faults via replication 
- match distribution of physical devices e.g. sensors
- increase security via isolation

### Main Topics
- Storage
- Communication
- Computation

### Map Reduce
- programming model for processing and generating large data sets
- [map] func for key/value pair to genreate intermediate k/v pairs
- [reduce] func for merge all intermediate values w/ same intermediate key
- functinoal styles programs are automatically parallelized and executed
- example : google tasks like indexing, data mining and machine learning 
- : count word frequency in documents, map: emit word with count, reduce: sums up counts for each word


## types
- map(k1, v1) -> list(k2, v2)
- reduce(k2, list(v2)) -> list(v2)

## Examples
- count of URL access frequency
- reverse web-link graph <target, sources> (target, list(source))
- term-vector per Host

## Applications
- google for indexing, processing logs, graph computations and machine learning
- more easy for large-scale distibuted systems

## Explanation
- Map: assign each friends (computer) few books to count the words
- Reduce: combine results after every friend finished