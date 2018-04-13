# strings-index

Try: $ go test -bench .

I'm getting (using go1.10 darwin/amd64):  
BenchmarkIndex1-8               300000000                4.01 ns/op  
BenchmarkIndex2-8               300000000                4.39 ns/op  
BenchmarkIndexByte1-8           1000000000               2.23 ns/op  
BenchmarkIndexByte2-8           1000000000               2.25 ns/op  
