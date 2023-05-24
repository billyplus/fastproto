package test

import (
	reflect "reflect"
	"testing"

	"github.com/billyplus/fastproto"
	"google.golang.org/protobuf/proto"
)

func TestMarshaler(t *testing.T) {
	testcase := getTestCases()

	for _, v := range testcase {
		testMarshaler(t, v)
	}

}

func testMarshaler(t *testing.T, v fastproto.Message) {
	data, err := fastproto.Marshal(v)
	if err != nil {
		t.Fatalf("marshal %T to []byte has error: %v", v, err)
	}

	v2 := reflect.New(reflect.TypeOf(v).Elem()).Interface().(proto.Message)
	err = proto.Unmarshal(data, v2)
	if err != nil {
		t.Fatalf("unmarshal %T with google proto has error: %v", v, err)
	}

	fastproto.IgnoreMessageInfo = false
	v3 := reflect.New(reflect.TypeOf(v).Elem()).Interface().(fastproto.Message)
	err = fastproto.Unmarshal(data, v3)
	if err != nil {
		t.Fatalf("unmarshal %T with fastproto option `IgnoreMessageInfo = false` has error: %v", v, err)
	}
	if !proto.Equal(v3, v2) {
		t.Fatalf("message[%T] with messageinfo should be equal to message from proto.Unmarshal", v)
		return
	}

	fastproto.IgnoreMessageInfo = true
	v4 := reflect.New(reflect.TypeOf(v).Elem()).Interface().(fastproto.Message)
	err = fastproto.Unmarshal(data, v4)

	if err != nil {
		t.Fatalf("unmarshal %T with fastproto option `IgnoreMessageInfo = true` has error: %v", v, err)
	}
	// ensure size is cached
	v4.Size()

	if !proto.Equal(v, v4) {
		t.Fatalf("message[%T] without messageinfo should be equal to original message", v)
		return
	}
}

func BenchmarkFastMarshalStringSlice(b *testing.B) {
	benchFastMarshal(b, arrStringTest)
}

func BenchmarkGoogleMarshalStringSlice(b *testing.B) {
	benchGoogleMarshal(b, arrStringTest)
}

func BenchmarkFastMarshalBytesSlice(b *testing.B) {
	benchFastMarshal(b, arrBytesTest)
}

func BenchmarkGoogleMarshalBytesSlice(b *testing.B) {
	benchGoogleMarshal(b, arrBytesTest)
}

func BenchmarkFastMarshalInt32Slice(b *testing.B) {
	benchFastMarshal(b, arrInt32Test)
}

func BenchmarkGoogleMarshalInt32Slice(b *testing.B) {
	benchGoogleMarshal(b, arrInt32Test)
}

func BenchmarkFastMarshalSint64Slice(b *testing.B) {
	benchFastMarshal(b, arrSint64Test)
}

func BenchmarkGoogleMarshalSint64Slice(b *testing.B) {
	benchGoogleMarshal(b, arrSint64Test)
}

func BenchmarkFastMarshalSfixed32Slice(b *testing.B) {
	benchFastMarshal(b, arrSfixed32Test)
}

func BenchmarkGoogleMarshalSfixed32Slice(b *testing.B) {
	benchGoogleMarshal(b, arrSfixed32Test)
}

func BenchmarkFastMarshalSfixed64Slice(b *testing.B) {
	benchFastMarshal(b, arrSfixed64Test)
}

func BenchmarkGoogleMarshalSfixed64Slice(b *testing.B) {
	benchGoogleMarshal(b, arrSfixed64Test)
}

func BenchmarkFastMarshalToMixedProto(b *testing.B) {
	benchFastMarshalTo(b, fullProtoTest)
}

func BenchmarkFastMarshalMixedProto(b *testing.B) {
	benchFastMarshal(b, fullProtoTest)
}

func BenchmarkGoogleMarshalMixedProto(b *testing.B) {
	benchGoogleMarshal(b, fullProtoTest)
}

func benchFastMarshalTo(b *testing.B, v fastproto.Marshaler) {
	d := make([]byte, 0, v.(fastproto.Sizer).Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fastproto.MarshalTo(d[:0], v)
	}
}

func benchFastMarshal(b *testing.B, v fastproto.Marshaler) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fastproto.Marshal(v)
	}
}

func benchGoogleMarshal(b *testing.B, v proto.Message) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(v)
	}
}
