package protohelper

import (
	"github.com/billyplus/fastproto"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

func ConsumeMessage(data []byte, msg proto.Message) (int, error) {
	msglen, n := CalcListLength(data)
	if n < 0 {
		return 0, protowire.ParseError(n)
	}
	data = data[n:]
	if err := fastproto.Unmarshal(data[:msglen], msg); err != nil {
		return 0, err
	}
	return n + msglen, nil
}
