# Protocol Buffers for Go with faster Marshal and Unmarshal methods

FastProto is a go support for adding extra `Marshal` and `Unmarshal` methods to standard generated protobuf file inspired by <a href="https://github.com/gogo/protobuf">gogo/protobuf</a>. For now, FastProto support proto 3 only.

## Getting Started

### Installation

Install the standard protocol buffer implementation from [https://github.com/protocolbuffers/protobuf](https://github.com/google/protobuf) first.

Then install the `protoc-gen-go-fast` binary

    go get github.com/billyplus/fastproto/cmd/protoc-gen-go-fast

### How to use

Generate standard `.pb.go` with protoc-gen-go first, then generate an extra `_fast.pb.go` file with `Marshal` and `Unmarshal` methods.

    protoc --go_out=./ ./test/msg.proto
    protoc --go-fast_out=./ ./test/msg.proto

### Marshal message

``` golang
    msg := &pb.SomeProtoMsg{}
    if err := fastproto.Marshal(msg); err!=nil {}

    // you can allocate []byte first
    data := make([]byte, msg.Size())
    if n, err := fastproto.MarshalTo(data, msg); err!=nil {}
    // result is data[:n]
```

### Unmarshal message

``` golang
    msg := &pb.SomeProtoMsg{}
    if err := fastproto.Unmarshal(data, &msg); err!=nil {}
```

### GRPC

It works with grpc.

Option 1. Replace the default codec for `proto`

``` golang
import "google.golang.org/grpc/encoding"

func main() {
    // replace the default codec.
	encoding.RegisterCodec(fastproto.ProtoCodec())
}
```

Option 2. **Not recommended**. Use `grpc.ForceServerCodec` option or `grpc.CustomCodec` option to create grpc server. Notice: This API is may be changed or removed in a later release.

``` golang
    import "google.golang.org/grpc"
    s := grpc.NewServer(grpc.CustomCodec(fastproto.ProtoCodec()))

    // or

    s := grpc.NewServer(grpc.ForceServerCodec(fastproto.ProtoCodec()))

```



