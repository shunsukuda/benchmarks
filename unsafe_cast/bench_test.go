package unsafecast_test

import (
	"reflect"
	"testing"
	"unsafe"
)

const (
	str = "abcdefghijklmnopqrstuvwxyz0123456789"
)

var (
	sstr = string(str)
	bstr = []byte(str)
)

func BenchmarkBytesToString(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = string(bstr)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = []byte(sstr)
	}
}

func BenchmarkUnsafeBytesToString(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = UnsafeBytesToString(bstr)
	}
}

func BenchmarkUnsafeStringToBytes(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = UnsafeStringToBytes(sstr)
	}
}
func BenchmarkUnsafeStringToBytesSlice(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = UnsafeStringToBytesSlice(sstr)
	}
}

func UnsafeBytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(unsafeBytesToStringHeader(&b)))
}

func UnsafeStringToBytesSlice(s string) []byte {
	bp := *(*[]byte)(unsafe.Pointer(unsafeStringToByteSliceHeader(&s)))
	return unsafe.Slice(&bp[0], len(s))
}

func UnsafeStringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(unsafeStringToByteSliceHeader(&s)))
}

func unsafeBytesToStringHeader(b *[]byte) *reflect.StringHeader {
	return &reflect.StringHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(b)).Data,
		Len:  len(*b),
	}
}

func unsafeStringToByteSliceHeader(s *string) *reflect.SliceHeader {
	return &reflect.SliceHeader{
		Data: (*reflect.StringHeader)(unsafe.Pointer(s)).Data,
		Len:  len(*s),
		Cap:  len(*s),
	}
}
