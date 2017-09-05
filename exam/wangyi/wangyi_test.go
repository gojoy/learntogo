package wangyi

import "testing"

func BenchmarkFashe(b *testing.B) {
	for i:=0;i<b.N;i++ {
		fashe(int64(b.N))
	}
}
