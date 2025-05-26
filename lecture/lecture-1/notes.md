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

## types
- map(k1, v1) -> list(k2, v2)
- reduce(k2, list(v2)) -> list(v2)

## Examples
- count of URL access frequency
- reverse web-link graph <target, sources> (target, list(source))
- term-vector per Host