package test

import (
	"testing"

	"github.com/billyplus/fastproto"
	"google.golang.org/protobuf/proto"
)

func TestSizer(t *testing.T) {
	testcase := getTestCases()

	for _, v := range testcase {
		testSizer(t, v.(fastproto.Sizer))
	}

}

func TestFullProtoSizer(t *testing.T) {
	testcase := fullProtoMsg()
	// testcase := &Uint64S{Val: randIntArr[uint64](20)}
	// testcase := &FullProto{ArrUint64: randIntArr[uint64](20)}

	testSizer(t, testcase)

}

func testSizer(t *testing.T, v fastproto.Sizer) {
	data, err := fastproto.Marshal(v.(fastproto.Message))
	if err != nil {
		t.Fatalf("marshal %T with fastproto has error: %v", v, err)
	}

	if len(data) != v.Size() {
		t.Fatalf("message %T size should equal len(data): %v", v, err)
	}
}

func BenchmarkFastSizeMixedProto(b *testing.B) {
	benchFastSizer(b, fullProtoTest)
}

func BenchmarkGoogleSizeMixedProto(b *testing.B) {
	benchGoogleSizer(b, fullProtoTest)
}

func benchFastSizer(b *testing.B, v fastproto.Sizer) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fastproto.Size(v)
	}
}

func benchGoogleSizer(b *testing.B, v proto.Message) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = proto.Size(v)
	}
}
