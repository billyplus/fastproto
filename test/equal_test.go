package test

import (
	"testing"

	"github.com/billyplus/fastproto"
	"google.golang.org/protobuf/proto"
)

// func TestNil(t *testing.T) {
// 	v1 := fullProtoMsg()
// 	v1.MActor = nil
// 	var v2 *FullProto
// 	v2 = nil
// 	checknil(v2, v1.MActor)
// 	// checknil(v1.MActor)
// 	// checknil(nil)
// 	// t.Fatal("finish")
// }

// func checknil(m1, m2 proto.Message) {
// 	if m1 == m2 {
// 		fmt.Println("is nil")
// 	} else {
// 		fmt.Println("not nil")
// 	}
// 	if m1 == nil {
// 		fmt.Println("m1 is nil")
// 	}
// 	if m2 == nil {
// 		fmt.Println("m2 is nil")
// 	}

// }

func TestEqual(t *testing.T) {
	v1 := fullProtoMsg()
	v2 := proto.Clone(v1)
	if !fastproto.Equal(v1, v2) {
		t.Fatalf("message %T should be equal", v1)
	}

	v3 := proto.Clone(v1).(*FullProto)
	v3.SInt32 += 100
	if fastproto.Equal(v1, v3) {
		t.Fatalf("message %T should not be equal", v1)
	}

	v3 = proto.Clone(v1).(*FullProto)
	v3.ArrInt32[len(v3.ArrInt32)-1] += 100
	if fastproto.Equal(v1, v3) {
		t.Fatalf("message %T should not be equal", v1)
	}

	v3 = proto.Clone(v1).(*FullProto)
	v3.ArrString[len(v3.ArrString)-1] = "test"
	if fastproto.Equal(v1, v3) {
		t.Fatalf("message %T should not be equal", v1)
	}

	v3 = proto.Clone(v1).(*FullProto)
	v3.MActor.Job += 100
	if fastproto.Equal(v1, v3) {
		t.Fatalf("message %T should not be equal", v1)
	}

	v3 = proto.Clone(v1).(*FullProto)
	v3.MActor.OpenId = "100"
	if fastproto.Equal(v1, v3) {
		t.Fatalf("message %T should not be equal", v1)
	}

	v3 = proto.Clone(v1).(*FullProto)
	v3.MActor = nil
	if fastproto.Equal(v1, v3) {
		t.Fatalf("message %T should not be equal", v1)
	}

	v3 = proto.Clone(v1).(*FullProto)
	v3.MapSfixed32Fixed64[100] = 334343
	if fastproto.Equal(v1, v3) {
		t.Fatalf("message %T should not be equal", v1)
	}

	v3 = proto.Clone(v1).(*FullProto)
	for k := range v3.MapStringActor {
		v3.MapStringActor[k] = nil
	}
	if fastproto.Equal(v1, v3) {
		t.Fatalf("message %T should not be equal", v1)
	}

	v3 = proto.Clone(v1).(*FullProto)
	for k := range v3.MapStringActor {
		v3.MapStringActor[k].OpenId = "nil"
	}
	if fastproto.Equal(v1, v3) {
		t.Fatalf("message %T should not be equal", v1)
	}
}

func BenchmarkFastEqual(b *testing.B) {
	v1 := fullProtoTest
	v2 := proto.Clone(v1)
	for i := 0; i < b.N; i++ {
		_ = fastproto.Equal(v1, v2)
	}
}

func BenchmarkGoogleEqual(b *testing.B) {
	v1 := fullProtoTest
	v2 := proto.Clone(v1)
	for i := 0; i < b.N; i++ {
		_ = proto.Equal(v1, v2)
	}
}
