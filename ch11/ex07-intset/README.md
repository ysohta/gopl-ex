
```
$ go test -bench=.
BenchmarkIntSetHas-4                	2000000000	         1.14 ns/op
BenchmarkIntSet32Has-4              	2000000000	         1.16 ns/op
BenchmarkIntSetBuiltinHas-4         	100000000	        13.9 ns/op
BenchmarkIntAdd-4                   	300000000	         5.49 ns/op
BenchmarkIntAdd32-4                 	300000000	         5.21 ns/op
BenchmarkIntSetBuiltinAdd-4         	50000000	        28.7 ns/op
BenchmarkIntUnionWith-4             	  200000	      7153 ns/op
BenchmarkIntUnionWith32-4           	  100000	     13284 ns/op
BenchmarkIntSetBuiltinUnionWith-4   	 3000000	       387 ns/op
PASS
ok  	github.com/ysohta/gopl-ex/ch11/ex07-intset	16.566s
```
