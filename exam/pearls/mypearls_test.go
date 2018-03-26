package pearls

import "testing"

func BenchmarkIsRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandIntNum(100, 10000)
	}
}

func BenchmarkRandIntByMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandIntByMap(100, 10000)
	}
}
