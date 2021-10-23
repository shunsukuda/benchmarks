package main

import "testing"

type MyInterface interface {
	Int() int
}

type MyInt int

func (e MyInt) Int() int {
	return int(e)
}

func (e MyInt) Int2() int {
	return int(e)
}

func BenchmarkCallInterfaceMethod(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var v MyInterface
	for i := 0; i < b.N; i++ {
		v = MyInt(i)
		v.(MyInt).Int2()
	}
}

func BenchmarkCallStructMethod(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var v MyInt
	for i := 0; i < b.N; i++ {
		v = MyInt(i)
		v.Int2()
	}
}
