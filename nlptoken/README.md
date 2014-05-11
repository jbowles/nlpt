#nlptoken

## Running Basic commands
* Run basic benchmarks
  * `go test -bench .`
  * `go test -run=none -bench=BenchmarkLexTknzBadStr -cpuprofile=cprof` writes to cprof.

* Run basic benchmarks with memory
  * `go test -benchmem -bench .`



## Running with pprof tools
* Run basic benchmarks
  * `go test -bench .`
  * `go test -bench . -cpuprofile=cprof` writes to cprof

## Benchmarking and Profiling
Without interface `Tokenizer` and tokenizer Digests:

```sh
BenchmarkUnicTokenGoodStr	        100000	      18441 ns/op	  12664 B/op	  187 allocs/op
BenchmarkUnicTokenBucketBadStr	  200000	      12400 ns/op	  7249 B/op	    133 allocs/op
BenchmarkWhiteSpace	              2000000000	  0.00 ns/op	  0 B/op	      0 allocs/op
```

With interface `Tokenizer` and Digests:

```sh
BenchmarkBuktTknzGoodStr	  100000	      18553 ns/op	    12808 B/op	    188 allocs/op
BenchmarkBuktTnkzBadStr	    200000	      12732 ns/op	    7374 B/op	      134 allocs/op
BenchmarkWhiteSpace	        2000000000	  0.00 ns/op	    0 B/op	        0 allocs/op

# new tokenizer
BenchmarkLexTknzGoodStr	    50000	        54472 ns/op	    6472 B/op	      178 allocs/op
BenchmarkLexTknzBadStr	    50000	        47198 ns/op	    6201 B/op	      167 allocs/op
```
