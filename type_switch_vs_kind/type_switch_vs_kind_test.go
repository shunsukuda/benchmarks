package main

import (
	"reflect"
	"testing"

	goreflect "github.com/goccy/go-reflect"
)

var (
	typesSlice = []interface{}{
		bool(false),
		int64(0),
		uint64(0),
		float64(0),
		string(""),
		&struct {
			I int64
			F float64
		}{0, 0},
	}
)

func BenchmarkTypeSwitch(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	l := len(typesSlice)
	for i := 0; i < b.N; i++ {
		typeswitch(typesSlice[i%l])
	}
}

func BenchmarkTypeKindT(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	l := len(typesSlice)
	for i := 0; i < b.N; i++ {
		typekindT(typesSlice[i%l])
	}
}

func BenchmarkTypeKindV(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	l := len(typesSlice)
	for i := 0; i < b.N; i++ {
		typekindV(typesSlice[i%l])
	}
}

func BenchmarkGoTypeKindT(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	l := len(typesSlice)
	for i := 0; i < b.N; i++ {
		gotypekindT(typesSlice[i%l])
	}
}

func BenchmarkGoTypeKindV(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	l := len(typesSlice)
	for i := 0; i < b.N; i++ {
		gotypekindV(typesSlice[i%l])
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

func typekindT(i interface{}) {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Bool:
		return
	case reflect.Int64:
		return
	case reflect.Uint64:
		return
	case reflect.Float64:
		return
	case reflect.String:
		return
	default:
		return
	}
}

func typekindV(i interface{}) {
	switch reflect.ValueOf(i).Kind() {
	case reflect.Bool:
		return
	case reflect.Int64:
		return
	case reflect.Uint64:
		return
	case reflect.Float64:
		return
	case reflect.String:
		return
	default:
		return
	}
}

func gotypekindT(i interface{}) {
	switch goreflect.TypeOf(i).Kind() {
	case goreflect.Bool:
		return
	case goreflect.Int64:
		return
	case goreflect.Uint64:
		return
	case goreflect.Float64:
		return
	case goreflect.String:
		return
	default:
		return
	}
}

func gotypekindV(i interface{}) {
	switch goreflect.ValueOf(i).Kind() {
	case goreflect.Bool:
		return
	case goreflect.Int64:
		return
	case goreflect.Uint64:
		return
	case goreflect.Float64:
		return
	case goreflect.String:
		return
	default:
		return
	}
}
