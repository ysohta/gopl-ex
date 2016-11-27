```bash
$ ./ex05-mandelbrot -help
Usage of ./ex05-mandelbrot:
  -n int
    	number of goroutine (default 4)

$ go test -bench .
testing: warning: no tests to run
PASS
BenchmarkGoroutine1-4   	       2	 594500140 ns/op
BenchmarkGoroutine2-4   	       3	 373401387 ns/op
BenchmarkGoroutine4-4   	       3	 368035074 ns/op
BenchmarkGoroutine8-4   	       3	 353785623 ns/op
BenchmarkGoroutine16-4  	       3	 350403226 ns/op
BenchmarkGoroutine32-4  	       3	 338974182 ns/op
BenchmarkGoroutine64-4  	       3	 334765551 ns/op
BenchmarkGoroutine1024-4	       3	 357236727 ns/op
ok  	github.com/ysohta/gopl-ex/ch08/ex05-mandelbrot	16.716s
```