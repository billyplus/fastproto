package test

import (
	"reflect"
	"testing"

	"github.com/billyplus/fastproto"
	"google.golang.org/protobuf/proto"
)

func TestUnmarshal(t *testing.T) {
	for _, v := range allTests {
		if testUnmarshal(t, v) != nil {
			return
		}
	}

}

func testUnmarshal(t *testing.T, v proto.Message) error {
	data, err := proto.Marshal(v)
	if err != nil {
		t.Fatalf("marshal %T to []byte has error: %v", v, err)
	}

	v2 := reflect.New(reflect.TypeOf(v).Elem()).Interface().(proto.Message)
	// err = protobuf.Unmarshal(data, v2)

	err = fastproto.Unmarshal(data, v2)
	if err != nil {
		t.Fatalf("unmarshal %T with fastproto has error: %v", v, err)
	}

	v3 := reflect.New(reflect.TypeOf(v).Elem()).Interface().(proto.Message)
	err = proto.Unmarshal(data, v3)
	if err != nil {
		t.Fatalf("unmarshal %T with google proto has error: %v", v, err)
	}

	// compared with value from google proto package
	if !proto.Equal(v3, v2.(fastproto.Message)) {
		t.Fatalf("message[%T] with messageinfo should be equal to message from proto.Unmarshal", v)
	}
	return nil
}

func BenchmarkFastUnmarshalStringSlice(b *testing.B) {
	benchFastUnmarshal(b, arrStringTest)
}

func BenchmarkGoogleUnmarshalStringSlice(b *testing.B) {
	benchGoogleUnmarshal(b, arrStringTest)
}

func BenchmarkFastUnmarshalBytesSlice(b *testing.B) {
	benchFastUnmarshal(b, arrBytesTest)
}

func BenchmarkGoogleUnmarshalBytesSlice(b *testing.B) {
	benchGoogleUnmarshal(b, arrBytesTest)
}

func BenchmarkFastUnmarshalInt32Slice(b *testing.B) {
	benchFastUnmarshal(b, arrInt32Test)
}

func BenchmarkGoogleUnmarshalInt32Slice(b *testing.B) {
	benchGoogleUnmarshal(b, arrInt32Test)
}

func BenchmarkFastUnmarshalSint64Slice(b *testing.B) {
	benchFastUnmarshal(b, arrSint64Test)
}

func BenchmarkGoogleUnmarshalSint64Slice(b *testing.B) {
	benchGoogleUnmarshal(b, arrSint64Test)
}

func BenchmarkFastUnmarshalSfixed32Slice(b *testing.B) {
	benchFastUnmarshal(b, arrSfixed32Test)
}

func BenchmarkGoogleUnmarshalSfixed32Slice(b *testing.B) {
	benchGoogleUnmarshal(b, arrSfixed32Test)
}

func BenchmarkFastUnmarshalSfixed64Slice(b *testing.B) {
	benchFastUnmarshal(b, arrSfixed64Test)
}

func BenchmarkGoogleUnmarshalSfixed64Slice(b *testing.B) {
	benchGoogleUnmarshal(b, arrSfixed64Test)
}

func BenchmarkFastUnmarshalMixedProto(b *testing.B) {
	benchFastUnmarshal(b, fullProtoTest)
}

func BenchmarkGoogleUnmarshalMixedProto(b *testing.B) {
	benchGoogleUnmarshal(b, fullProtoTest)
}

func benchFastUnmarshal(b *testing.B, v proto.Message) {
	data, _ := proto.Marshal(v)
	t := reflect.TypeOf(v).Elem()
	v2 := reflect.New(t).Interface().(proto.Message)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// b.StopTimer()
		// b.StartTimer()
		_ = fastproto.Unmarshal(data, v2)
	}
}

func benchGoogleUnmarshal(b *testing.B, v proto.Message) {
	data, _ := proto.Marshal(v)
	t := reflect.TypeOf(v).Elem()
	v3 := reflect.New(t).Interface().(proto.Message)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// b.StopTimer()
		// b.StartTimer()
		_ = proto.Unmarshal(data, v3)

	}
}
