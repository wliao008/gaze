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

## Benchmarks

Benchmarks are ran on the following hardwares:

1. Intel(R) Core(TM)2 Quad CPU    Q6600  @ 2.40GHz
2. Macbook Pro

Hardware | Algorithm | Size | Iterations | ns/ops | ms/op
-------------|------------------|-------|------|-----------|----------
1| Backtracking | 50x25 | 50 | 22865910 |22.86 
1| Backtracking | 100x50 | 20 | 91300763 | 91.30
1| Backtracking | 1000x500 | 1 | 9776425846 | 9776.42 
1| Kruskal | 50x25 | 2000 | 783913 | 0.78
1| Kruskal | 100x50 | 500 | 3147580 | 3.14
1| Kruskal | 1000x500 | 3 | 3147580 | 363.81
