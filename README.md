# bits

bits is a bitwise operation library that does bulk operations on architectures that support unaligned read/writes.

Based on go source code https://golang.org/src/crypto/cipher/xor.go

Benchmark results for bulk/safe Xor on a i5 2500k

~~~
goos: windows
goarch: amd64
pkg: github.com/serdaroquai/bits
BenchmarkFastXor-4   	200000000	         8.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkSafeXor-4   	100000000	        15.9 ns/op	       0 B/op	       0 allocs/op
~~~