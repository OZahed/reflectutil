# [WIP] reflectutil

is a repository for casting Goalng data Types regarding struct field names and field types

This repository is still working in progress

based on benchmarks below which you can find the actual benchmark functions for common type casts in `test/commoncast` package using reflectutil is

benchmark:

```bash
go test -bench=. -benchmem  -count=5 ./test/...
```

how to read benchmark results:
with this set of keys the output would look like this

| Benchmark-name-CPU_COUNT             | number of times function have been called | how long it took for function call | number of bytes for that operation | number of times memory allocation happened |
| :----------------------------------- | :---------------------------------------: | :--------------------------------: | :--------------------------------: | :----------------------------------------: |
| BenchmarkTypeCast-12                 |                   19533                   |            64323 ns/op             |             4832 B/op              |               168 allocs/op                |
| BenchmarkTypeCast-12                 |                   17467                   |            62523 ns/op             |             4832 B/op              |               168 allocs/op                |
| BenchmarkTypeCast-12                 |                   18489                   |            63172 ns/op             |             4832 B/op              |               168 allocs/op                |
| BenchmarkTypeCast-12                 |                   18308                   |            63614 ns/op             |             4832 B/op              |               168 allocs/op                |
| BenchmarkTypeCast-12                 |                   17440                   |            62514 ns/op             |             4832 B/op              |               168 allocs/op                |
| ---                                  |                 ---------                 |                ---                 |                ---                 |                -----------                 |
| BenchmarkCommonCast-12               |                   46586                   |            25790 ns/op             |             1016 B/op              |                14 allocs/op                |
| BenchmarkCommonCast-12               |                   45694                   |            25986 ns/op             |             1016 B/op              |                14 allocs/op                |
| BenchmarkCommonCast-12               |                   47884                   |            25715 ns/op             |             1016 B/op              |                14 allocs/op                |
| BenchmarkCommonCast-12               |                   48811                   |            26662 ns/op             |             1016 B/op              |                14 allocs/op                |
| BenchmarkCommonCast-12               |                   44476                   |            25286 ns/op             |             1016 B/op              |                14 allocs/op                |
| ----                                 |                ----------                 |                ---                 |                ---                 |                -----------                 |
| BenchmarkCommonCastArray-12          |                   44493                   |            27283 ns/op             |             2536 B/op              |                26 allocs/op                |
| BenchmarkCommonCastArray-12          |                   42986                   |            27993 ns/op             |             2536 B/op              |                26 allocs/op                |
| BenchmarkCommonCastArray-12          |                   43272                   |            28413 ns/op             |             2536 B/op              |                26 allocs/op                |
| BenchmarkCommonCastArray-12          |                   39248                   |            28329 ns/op             |             2536 B/op              |                26 allocs/op                |
| BenchmarkCommonCastArray-12          |                   43066                   |            27778 ns/op             |             2536 B/op              |                26 allocs/op                |
| ----                                 |                ----------                 |                ---                 |                ---                 |                -----------                 |
| BenchmarkCastCommonSliceAppend-12    |                   42824                   |            29113 ns/op             |             2696 B/op              |                27 allocs/op                |
| BenchmarkCastCommonSliceAppend-12    |                   34981                   |            28753 ns/op             |             2696 B/op              |                27 allocs/op                |
| BenchmarkCastCommonSliceAppend-12    |                   38353                   |            28625 ns/op             |             2696 B/op              |                27 allocs/op                |
| BenchmarkCastCommonSliceAppend-12    |                   43855                   |            28401 ns/op             |             2696 B/op              |                27 allocs/op                |
| BenchmarkCastCommonSliceAppend-12    |                   43156                   |            32718 ns/op             |             2696 B/op              |                27 allocs/op                |
| ----                                 |                ----------                 |                ---                 |                ---                 |                -----------                 |
| BenchmarkTypeCastHandWrittenLoop-12  |                   10000                   |            104527 ns/op            |             9848 B/op              |               332 allocs/op                |
| BenchmarkTypeCastHandWrittenLoop-12  |                   8641                    |            134107 ns/op            |             9848 B/op              |               332 allocs/op                |
| BenchmarkTypeCastHandWrittenLoop-12  |                   8577                    |            124294 ns/op            |             9848 B/op              |               332 allocs/op                |
| BenchmarkTypeCastHandWrittenLoop-12  |                   9202                    |            125424 ns/op            |             9848 B/op              |               332 allocs/op                |
| BenchmarkTypeCastHandWrittenLoop-12  |                   9702                    |            123243 ns/op            |             9848 B/op              |               332 allocs/op                |
| ----                                 |                ----------                 |                ---                 |                ---                 |                -----------                 |
| BenchmarkTypeCastArrayReflectUtil-12 |                   8230                    |            133717 ns/op            |             10544 B/op             |               361 allocs/op                |
| BenchmarkTypeCastArrayReflectUtil-12 |                   8632                    |            132978 ns/op            |             10544 B/op             |               361 allocs/op                |
| BenchmarkTypeCastArrayReflectUtil-12 |                   8594                    |            137943 ns/op            |             10544 B/op             |               361 allocs/op                |
| BenchmarkTypeCastArrayReflectUtil-12 |                   7273                    |            137575 ns/op            |             10544 B/op             |               361 allocs/op                |
| BenchmarkTypeCastArrayReflectUtil-12 |                   7872                    |            138412 ns/op            |             10544 B/op             |               361 allocs/op                |
