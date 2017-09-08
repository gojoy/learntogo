package jd

import "testing"

func BenchmarkQ2(b *testing.B) {
	for i:=0;i<b.N;i++ {
		Q2()
	}
}
