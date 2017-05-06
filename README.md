## Usage

```
$ go run main.go
   _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
| |  _ _ _ _ _ _ _ _ _ _  |  _ _ _|  _ _ _ _ _    |_   _  |_   _ _ _|  _ _ _ _  |
| |_ _ _ _ _ _ _  |_   _| |_ _ _ _ _ _|  _ _ _ _| |  _|  _|  _|  _ _ _ _ _ _|   |
|_ _ _ _ _ _ _  |_  |_ _ _ _ _ _  |  _ _|_ _ _ _ _ _| |_  |_ _ _ _ _|  _ _  | | |
|  _ _ _  |   |_  |_ _  |  _ _ _ _|_   _ _ _ _  |_ _ _  | |  _  |  _ _|  _ _ _|_|
|_ _ _  |_ _|_ _|_ _  | |_ _ _  |   |_ _ _   _|_   _  |_ _ _|  _|_ _  |_ _|  _  |
|    _ _|  _ _ _ _ _ _|  _ _  | | |_ _  |_ _|   |_  |_ _ _  |_ _ _  | |  _ _|  _|
| |_|  _ _|  _ _  |  _ _|  _ _|_ _|   |_ _ _ _| | |  _ _ _|_  |  _| | |_ _  |_  |
|   |_ _ _ _|  _ _|_  | |_ _ _  |  _| | |  _ _ _|_ _|  _ _ _ _|  _ _|_ _ _ _ _| |
| |_   _ _ _| |   |_ _ _ _|  _ _| |_ _ _| |_    |   |_ _ _ _  | |_   _ _ _ _  | |
|_ _|_ _ _ _ _ _|_ _ _ _ _ _|_ _ _ _ _ _ _ _ _|_ _|_ _ _ _ _ _|_ _ _|_ _ _ _ _  |
```

## Algorithm Benchmarks

Recursive backtracking is a pretty brute force algorithm, a maze of 50 x 25 takes about 25 milliseconds to generate. After some optimzation, it's reduced to roughly 12 milliseconds. 

50 x 25
```
$ go test -v ./... -bench=.
 BenchmarkBackTrackingAlgo-4          100          12557345 ns/op 
 BenchmarkKruskalAlgo-4              5000            275315 ns/op
```

100 x 50
```
$ go test -v ./... -bench=.
 BenchmarkBackTrackingAlgo-4           20          50495241 ns/op
 BenchmarkKruskalAlgo-4              2000           1101793 ns/op
```
