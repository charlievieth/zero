package zero_test

import (
	"bytes"
	"testing"

	"github.com/charlievieth/zero"
)

const bufSize = 1024 * 1024

var zeroBuf = make([]byte, bufSize)
var oneBuf = make([]byte, bufSize)

func init() {
	for i := 0; i < len(zeroBuf); i++ {
		zeroBuf[i] = 0
	}
	for i := 0; i < len(oneBuf); i++ {
		oneBuf[i] = 0
	}
	oneBuf[len(oneBuf)-1] = 1
}

func BenchmarkZero_One_128k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if zero.Zero(oneBuf) {
			b.Fatal("WTF")
		}
	}
}

func BenchmarkEqual_One_128k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if bytes.Equal(oneBuf, zeroBuf) {
			b.Fatal("WTF")
		}
	}
}

func slowZero(b []byte) bool {
	for _, c := range b {
		if c != 0 {
			return false
		}
	}
	return true
}

func BenchmarkSlow_One_128k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if slowZero(oneBuf) {
			b.Fatal("WTF")
		}
	}
}
