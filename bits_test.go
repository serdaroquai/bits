package bits

import (
	"bytes"
	"testing"
)

func TestXor(t *testing.T) {

	for alignP := 0; alignP < 2; alignP++ {
		for alignQ := 0; alignQ < 2; alignQ++ {
			for alignD := 0; alignD < 2; alignD++ {
				p := make([]byte, 1024)[alignP:]
				q := make([]byte, 1024)[alignQ:]
				d1 := make([]byte, 1024+alignD)[alignD:]
				d2 := make([]byte, 1024+alignD)[alignD:]
				fastXor(d1, p, q)
				safeXor(d2, p, q)
				if !bytes.Equal(d1, d2) {
					t.Error("not equal")
				}
			}
		}
	}
}

func TestAnd(t *testing.T) {

	for alignP := 0; alignP < 2; alignP++ {
		for alignQ := 0; alignQ < 2; alignQ++ {
			for alignD := 0; alignD < 2; alignD++ {
				p := make([]byte, 1024)[alignP:]
				q := make([]byte, 1024)[alignQ:]
				d1 := make([]byte, 1024+alignD)[alignD:]
				d2 := make([]byte, 1024+alignD)[alignD:]
				fastAnd(d1, p, q)
				safeAnd(d2, p, q)
				if !bytes.Equal(d1, d2) {
					t.Error("not equal")
				}
			}
		}
	}
}

func TestOr(t *testing.T) {

	for alignP := 0; alignP < 2; alignP++ {
		for alignQ := 0; alignQ < 2; alignQ++ {
			for alignD := 0; alignD < 2; alignD++ {
				p := make([]byte, 1024)[alignP:]
				q := make([]byte, 1024)[alignQ:]
				d1 := make([]byte, 1024+alignD)[alignD:]
				d2 := make([]byte, 1024+alignD)[alignD:]
				fastOr(d1, p, q)
				safeOr(d2, p, q)
				if !bytes.Equal(d1, d2) {
					t.Error("not equal")
				}
			}
		}
	}
}

var result int

func BenchmarkFastXor(b *testing.B) {
	var r int
	k := make([]byte, 8)
	l := make([]byte, 8)
	dst := make([]byte, 8)

	l[0] = 0xFF
	for i := 1; i < 8; i++ {
		l[i] = l[i-1] >> 1
	}

	for n := 0; n < b.N; n++ {

		// always record the result to prevent
		// the compiler eliminating the function call.
		r = fastXor(dst, k, l)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkSafeXor(b *testing.B) {
	var r int
	k := make([]byte, 8)
	l := make([]byte, 8)
	dst := make([]byte, 8)

	l[0] = 0xFF
	for i := 1; i < 8; i++ {
		l[i] = l[i-1] >> 1
	}

	for n := 0; n < b.N; n++ {

		// always record the result to prevent
		// the compiler eliminating the function call.
		r = safeXor(dst, k, l)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
