GOMAXPROCS=4  is the best performance in my environment.

$ GOMAXPROCS=1 go test -bench .
BenchmarkGoroutine1   	       2	 579429382 ns/op
BenchmarkGoroutine2   	       2	 602183533 ns/op
BenchmarkGoroutine4   	       2	 598286484 ns/op
BenchmarkGoroutine8   	       2	 581357044 ns/op
BenchmarkGoroutine16  	       2	 582582629 ns/op
BenchmarkGoroutine32  	       2	 600817925 ns/op
BenchmarkGoroutine64  	       2	 601840377 ns/op
BenchmarkGoroutine1024	       2	 564071310 ns/op

$ GOMAXPROCS=2 go test -bench .
BenchmarkGoroutine1-2   	       2	 572799374 ns/op
BenchmarkGoroutine2-2   	       3	 381973380 ns/op
BenchmarkGoroutine4-2   	       3	 398898349 ns/op
BenchmarkGoroutine8-2   	       3	 398847939 ns/op
BenchmarkGoroutine16-2  	       3	 404993053 ns/op
BenchmarkGoroutine32-2  	       3	 396320993 ns/op
BenchmarkGoroutine64-2  	       3	 414689640 ns/op
BenchmarkGoroutine1024-2	       3	 410949334 ns/op

$ GOMAXPROCS=4 go test -bench .
BenchmarkGoroutine1-4   	       2	 551468831 ns/op
BenchmarkGoroutine2-4   	       3	 361430942 ns/op
BenchmarkGoroutine4-4   	       3	 353460552 ns/op
BenchmarkGoroutine8-4   	       3	 362103604 ns/op
BenchmarkGoroutine16-4  	       3	 346328316 ns/op
BenchmarkGoroutine32-4  	       3	 356437145 ns/op
BenchmarkGoroutine64-4  	       3	 358002006 ns/op
BenchmarkGoroutine1024-4	       3	 360626717 ns/op

$ GOMAXPROCS=8 go test -bench .
BenchmarkGoroutine1-8   	       2	 568508022 ns/op
BenchmarkGoroutine2-8   	       3	 378984650 ns/op
BenchmarkGoroutine4-8   	       3	 364283335 ns/op
BenchmarkGoroutine8-8   	       3	 371191315 ns/op
BenchmarkGoroutine16-8  	       3	 370951516 ns/op
BenchmarkGoroutine32-8  	       3	 354450390 ns/op
BenchmarkGoroutine64-8  	       3	 354625559 ns/op
BenchmarkGoroutine1024-8	       3	 397196786 ns/op

$ GOMAXPROCS=16 go test -bench .
BenchmarkGoroutine1-16   	       2	 565299295 ns/op
BenchmarkGoroutine2-16   	       3	 363872890 ns/op
BenchmarkGoroutine4-16   	       3	 392145776 ns/op
BenchmarkGoroutine8-16   	       3	 391983637 ns/op
BenchmarkGoroutine16-16  	       3	 382515463 ns/op
BenchmarkGoroutine32-16  	       3	 369649290 ns/op
BenchmarkGoroutine64-16  	       3	 374465229 ns/op
BenchmarkGoroutine1024-16	       3	 416802111 ns/op

