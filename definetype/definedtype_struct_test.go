package main

import "testing"

type DefinedTypeInt int

type StructTypeInt struct {
	V int
}

func BenchmarkDefineTypeNative(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := i
		_ = v
	}
}

func BenchmarkDefineTypeNativeCast(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := int(i)
		_ = v
	}
}

func BenchmarkDefinedTypeCast(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := DefinedTypeInt(i)
		_ = v
	}
}

func BenchmarkDefineTypeStruct(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := StructTypeInt{i}
		_ = v
	}
}
