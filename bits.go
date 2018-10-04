// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Adapted from: https://golang.org/src/crypto/cipher/xor.go

package bits

import (
	"runtime"
	"unsafe"
)

const wordSize = int(unsafe.Sizeof(uintptr(0)))
const supportsUnaligned = runtime.GOARCH == "386" || runtime.GOARCH == "amd64" || runtime.GOARCH == "ppc64" || runtime.GOARCH == "ppc64le" || runtime.GOARCH == "s390x"

// Xor []byte a,b into dst, assume dst has sufficient size
func Xor(dst, a, b []byte) int {
	if supportsUnaligned {
		return fastXor(dst, a, b)
	} else {
		return safeXor(dst, a, b)
	}
}

// And []byte a,b into dst, assume dst has sufficient size
func And(dst, a, b []byte) int {
	if supportsUnaligned {
		return fastAnd(dst, a, b)
	} else {
		return safeAnd(dst, a, b)
	}
}

// Or []byte a,b into dst, assume dst has sufficient size
func Or(dst, a, b []byte) int {
	if supportsUnaligned {
		return fastOr(dst, a, b)
	} else {
		return safeOr(dst, a, b)
	}
}

func fastXor(dst, a, b []byte) (length int) {
	length = len(a)
	if len(b) < length {
		length = len(b)
	}

	wordCount := length / wordSize

	if wordCount > 0 {
		dw := *(*[]uintptr)(unsafe.Pointer(&dst))
		da := *(*[]uintptr)(unsafe.Pointer(&a))
		db := *(*[]uintptr)(unsafe.Pointer(&b))

		for i := 0; i < wordCount; i++ {
			dw[i] = da[i] ^ db[i]
		}
	}

	for i := length - length%wordSize; i < length; i++ {
		dst[i] = a[i] ^ b[i]
	}

	return
}

func safeXor(dst, a, b []byte) (length int) {
	length = len(a)
	if len(b) < length {
		length = len(b)
	}
	for i := 0; i < length; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return length
}

func fastAnd(dst, a, b []byte) (length int) {
	length = len(a)
	if len(b) < length {
		length = len(b)
	}

	wordCount := length / wordSize

	if wordCount > 0 {
		dw := *(*[]uintptr)(unsafe.Pointer(&dst))
		da := *(*[]uintptr)(unsafe.Pointer(&a))
		db := *(*[]uintptr)(unsafe.Pointer(&b))

		for i := 0; i < wordCount; i++ {
			dw[i] = da[i] & db[i]
		}
	}

	for i := length - length%wordSize; i < length; i++ {
		dst[i] = a[i] & b[i]
	}

	return
}

func safeAnd(dst, a, b []byte) (length int) {
	length = len(a)
	if len(b) < length {
		length = len(b)
	}
	for i := 0; i < length; i++ {
		dst[i] = a[i] & b[i]
	}
	return length
}

func fastOr(dst, a, b []byte) (length int) {
	length = len(a)
	if len(b) < length {
		length = len(b)
	}

	wordCount := length / wordSize

	if wordCount > 0 {
		dw := *(*[]uintptr)(unsafe.Pointer(&dst))
		da := *(*[]uintptr)(unsafe.Pointer(&a))
		db := *(*[]uintptr)(unsafe.Pointer(&b))

		for i := 0; i < wordCount; i++ {
			dw[i] = da[i] | db[i]
		}
	}

	for i := length - length%wordSize; i < length; i++ {
		dst[i] = a[i] | b[i]
	}

	return
}

func safeOr(dst, a, b []byte) (length int) {
	length = len(a)
	if len(b) < length {
		length = len(b)
	}
	for i := 0; i < length; i++ {
		dst[i] = a[i] | b[i]
	}
	return length
}
