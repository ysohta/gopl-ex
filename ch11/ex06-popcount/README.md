
```
$ go test -bench=.
BenchmarkPopCount-4              	2000000000	         0.37 ns/op
BenchmarkShiftPopCount-4         	30000000	        54.6 ns/op
BenchmarkPopCountClearMinBit-4   	300000000	         6.20 ns/op
PASS
ok  	github.com/ysohta/gopl-ex/ch11/ex06-popcount	4.943s
```
