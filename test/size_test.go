package test

import (
	"testing"

	"github.com/billyplus/fastproto"
	"google.golang.org/protobuf/proto"
)

func TestSizer(t *testing.T) {
	testcase := getTestCases()

	for _, v := range testcase {
		testSizer(t, v)
	}

}

func testSizer(t *testing.T, v proto.Message) {
	if vv, ok := v.(fastproto.Message); ok {
		data, err := fastproto.Marshal(vv)
		if err != nil {
			t.Fatalf("marshal %T with fastproto has error: %v", v, err)
		}

		if len(data) != vv.Size() {
			t.Fatalf("message %T size should equal len(data): %v", v, err)
		}
	}
}

func BenchmarkFastSizeMixedProto(b *testing.B) {
	benchFastSizer(b, fullProtoTest)
}

func BenchmarkGoogleSizeMixedProto(b *testing.B) {
	benchGoogleSizer(b, fullProtoTest)
}

func benchFastSizer(b *testing.B, v proto.Message) {
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
