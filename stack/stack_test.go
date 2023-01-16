package stack

import (
	s "github.com/golang-collections/collections/stack"
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

func BenchmarkPush(b *testing.B) {
	s := NewWithSize[int](100)
	// str := RandStringBytesMaskImprSrcUnsafe(128)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkStandartPush(b *testing.B) {
	s := s.New()

	// str := RandStringBytesMaskImprSrcUnsafe(128)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
