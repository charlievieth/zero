package zero_test

import (
	"bytes"
	"testing"

	"github.com/charlievieth/zero"
)

var zeroBuf = make([]byte, 128*1024)
var oneBuf = make([]byte, 128*1024)

func init() {
	for i := 0; i < len(zeroBuf); i++ {
		zeroBuf[i] = 0
	}
	for i := 0; i < len(oneBuf); i++ {
		oneBuf[i] = 0
	}
	oneBuf[len(oneBuf)-1] = 1
}

func BenchmarkEqual_One_128k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if bytes.Equal(oneBuf, zeroBuf) {
			b.Fatal("WTF")
		}
	}
}

func BenchmarkZero_One_128k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if zero.Zero(oneBuf) {
			b.Fatal("WTF")
		}
	}
}
