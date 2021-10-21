package main

import (
	"reflect"
	"testing"

	goreflect "github.com/goccy/go-reflect"
)

func BenchmarkNativeIsZero(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = nativeiszero(i % 10)
	}
}

func BenchmarkReflectIsZero(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = reflectiszero(i % 10)
	}
}

func BenchmarkGoReflectIsZero(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = goreflectiszero(i % 10)
	}
}

func typeswitch(i interface{}) {
	switch i.(type) {
	case bool:
		return
	case int64:
		return
	case uint64:
		return
	case float64:
		return
	case string:
		return
	default:
		return
	}
}

func nativeiszero(i interface{}) bool {
	return (i == 0)
}

func reflectiszero(i interface{}) bool {
	return reflect.ValueOf(i).IsZero()
}

func goreflectiszero(i interface{}) bool {
	return goreflect.ValueOf(i).IsZero()
}
