package test

import (
	"math/rand"
	"testing"

	"github.com/billyplus/fastproto"
)

func TestUnknownFields(t *testing.T) {
	// testcase := getTestCases()
	v := &OtherMessage{
		Eid:    rand.Int63(),
		OpenId: "\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E",
		Name:   randStringN(20)(),
		Job:    rand.Int31(),
		Sex:    rand.Int31(),
	}

	data, err := fastproto.Marshal(v)
	if err != nil {
		t.Fatalf("Marshal %T with fastproto has error: %v", v, err)
	}

	var mm LessOtherMessage
	err = fastproto.Unmarshal(data, &mm)
	if err != nil {
		t.Fatalf("Unmarshal %T with fastproto has error: %v", v, err)
	}

	// for _, v := range testcase {
	// 	testSizer(t, v)
	// }
}

// func testUnknownFields(t *testing.T, v proto.Message) {
// 	if vv, ok := v.(fastproto.Message); ok {
// 		data, err := fastproto.Marshal(vv)
// 		if err != nil {
// 			t.Fatalf("marshal %T with fastproto has error: %v", v, err)
// 		}

// 		if len(data) != vv.Size() {
// 			t.Fatalf("message %T size should equal len(data): %v", v, err)
// 		}
// 	}
// }
