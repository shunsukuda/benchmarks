package main

import "testing"

func BenchmarkIntCast(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var x int
	for i := 0; i < b.N; i++ {
		x = i
		_ = int(x)
	}
}

func BenchmarkInterfaceTypeAssertion(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var x interface{}
	for i := 0; i < b.N; i++ {
		x = i
		_ = x.(int)
	}
}
