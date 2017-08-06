# go-rand

Pooled, lock-free xorshift PRNG drop-in replacement for Go

## Benchmark

```
BenchmarkCoreGlobal-4             50000000       29.2 ns/op
BenchmarkRandGlobal-4            100000000       21.5 ns/op

BenchmarkCore-4                  200000000       8.90 ns/op
BenchmarkRand-4                 1000000000       2.88 ns/op

BenchmarkCoreEach-4                 200000      10298 ns/op
BenchmarkRandEach-4               30000000       46.1 ns/op

BenchmarkCoreGlobalParallel-4     10000000        132 ns/op
BenchmarkRandGlobalParallel-4     50000000       26.9 ns/op

BenchmarkCoreParallel-4         1000000000       2.74 ns/op
BenchmarkRandParallel-4         2000000000       1.01 ns/op

BenchmarkCoreEachParallel-4         500000       3151 ns/op
BenchmarkRandEachParallel-4      100000000       14.0 ns/op
```
