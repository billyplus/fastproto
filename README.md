# Protocol Buffers for Go with faster Marshal and Unmarshal methods

FastProto is a go support for adding extra `Marshal` and `Unmarshal` methods to standard generated protobuf file inspired by <a href="https://github.com/gogo/protobuf">gogo/protobuf</a>. For now, FastProto support proto 3 only.

## Getting Started

### Installation

Install the standard protocol buffer implementation from [https://github.com/protocolbuffers/protobuf](https://github.com/google/protobuf) first.

Then install the `protoc-gen-go-fast` binary

    go install github.com/billyplus/fastproto/cmd/protoc-gen-go-fast@latest

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

### Options to control code generation

if you want to ignore marshal method

```
    import "github.com/billyplus/fastproto/options/options.proto"

    // if true, Marshal interface will not be generated for this file
    option (options.fastproto_no_marshaler) = true;

    message XXXX {
        // if true, Marshal interface will be generated for this message whatever fastproto_no_marshaler is.
        option (options.fastproto_msg_marshaler) = true;
        // if true, Marshal interface will not be generated for this message
        option (options.fastproto_msg_no_marshaler) = true;
    }

```

if you want to ignore unmarshal method

```
    import "github.com/billyplus/fastproto/options/options.proto"

    // if true, Unmarshal interface will not be generated for this file
    option (options.fastproto_no_unmarshaler) = true;

    message XXXX {
        // if true, Unmarshal interface will be generated for this message whatever fastproto_no_unmarshaler is.
        option (options.fastproto_msg_unmarshaler) = true;
        // if true, Unmarshal interface will not be generated for this message
        option (options.fastproto_msg_no_unmarshaler) = true;
    }

```

please check `options/options.proto` for other options.

### Benchmarks(go-1.24)

[code](https://github.com/billyplus/fastproto/tree/main/test)

`go test -benchmem -bench=. ./test --count=1 -benchtime=2s`

```
goos: linux
goarch: amd64
pkg: github.com/billyplus/fastproto/test
cpu: AMD Ryzen 9 5950X 16-Core Processor
```
  |                                          |          |              |            |                |
  | ---------------------------------------- | -------- | ------------ | ---------- | -------------- |
  | BenchmarkFastEqual-12                    | 297753   | 8112 ns/op   | 0 B/op     | 0 allocs/op    |
  | BenchmarkGoogleEqual-12                  | 20384    | 121474 ns/op | 34187 B/op | 3196 allocs/op |
  | BenchmarkFastMarshalStringSlice-12       | 34819057 | 70.72 ns/op  | 80 B/op    | 1 allocs/op    |
  | BenchmarkGoogleMarshalStringSlice-12     | 15401488 | 146.6 ns/op  | 80 B/op    | 1 allocs/op    |
  | BenchmarkFastMarshalBytesSlice-12        | 30732194 | 76.50 ns/op  | 80 B/op    | 1 allocs/op    |
  | BenchmarkGoogleMarshalBytesSlice-12      | 21982063 | 107.4 ns/op  | 80 B/op    | 1 allocs/op    |
  | BenchmarkFastMarshalInt32Slice-12        | 12266313 | 195.5 ns/op  | 128 B/op   | 1 allocs/op    |
  | BenchmarkGoogleMarshalInt32Slice-12      | 10538212 | 229.3 ns/op  | 128 B/op   | 1 allocs/op    |
  | BenchmarkFastMarshalSint64Slice-12       | 9308516  | 255.1 ns/op  | 224 B/op   | 1 allocs/op    |
  | BenchmarkGoogleMarshalSint64Slice-12     | 8185350  | 297.0 ns/op  | 224 B/op   | 1 allocs/op    |
  | BenchmarkFastMarshalSfixed32Slice-12     | 47673926 | 49.92 ns/op  | 112 B/op   | 1 allocs/op    |
  | BenchmarkGoogleMarshalSfixed32Slice-12   | 26201978 | 90.12 ns/op  | 112 B/op   | 1 allocs/op    |
  | BenchmarkFastMarshalSfixed64Slice-12     | 32481122 | 71.66 ns/op  | 208 B/op   | 1 allocs/op    |
  | BenchmarkGoogleMarshalSfixed64Slice-12   | 22325599 | 107.6 ns/op  | 208 B/op   | 1 allocs/op    |
  | BenchmarkFastMarshalToMixedProto-12      | 172848   | 13498 ns/op  | 0 B/op     | 0 allocs/op    |
  | BenchmarkFastMarshalMixedProto-12        | 106030   | 23164 ns/op  | 18432 B/op | 1 allocs/op    |
  | BenchmarkGoogleMarshalMixedProto-12      | 30116    | 79646 ns/op  | 32544 B/op | 1481 allocs/op |
  | BenchmarkFastSizeMixedProto-12           | 438618   | 5723 ns/op   | 0 B/op     | 0 allocs/op    |
  | BenchmarkGoogleSizeMixedProto-12         | 79185    | 29972 ns/op  | 7056 B/op  | 740 allocs/op  |
  | BenchmarkFastUnmarshalStringSlice-12     | 9100161  | 264.9 ns/op  | 378 B/op   | 8 allocs/op    |
  | BenchmarkGoogleUnmarshalStringSlice-12   | 6291268  | 381.4 ns/op  | 378 B/op   | 8 allocs/op    |
  | BenchmarkFastUnmarshalBytesSlice-12      | 7002562  | 338.8 ns/op  | 512 B/op   | 9 allocs/op    |
  | BenchmarkGoogleUnmarshalBytesSlice-12    | 5809155  | 399.8 ns/op  | 512 B/op   | 9 allocs/op    |
  | BenchmarkFastUnmarshalInt32Slice-12      | 12064306 | 197.6 ns/op  | 176 B/op   | 2 allocs/op    |
  | BenchmarkGoogleUnmarshalInt32Slice-12    | 9365703  | 256.6 ns/op  | 176 B/op   | 2 allocs/op    |
  | BenchmarkFastUnmarshalSint64Slice-12     | 7890279  | 304.7 ns/op  | 272 B/op   | 2 allocs/op    |
  | BenchmarkGoogleUnmarshalSint64Slice-12   | 6766574  | 357.3 ns/op  | 272 B/op   | 2 allocs/op    |
  | BenchmarkFastUnmarshalSfixed32Slice-12   | 18561280 | 128.7 ns/op  | 176 B/op   | 2 allocs/op    |
  | BenchmarkGoogleUnmarshalSfixed32Slice-12 | 16476768 | 143.3 ns/op  | 176 B/op   | 2 allocs/op    |
  | BenchmarkFastUnmarshalSfixed64Slice-12   | 13286451 | 181.0 ns/op  | 272 B/op   | 2 allocs/op    |
  | BenchmarkGoogleUnmarshalSfixed64Slice-12 | 14421873 | 166.1 ns/op  | 272 B/op   | 2 allocs/op    |
  | BenchmarkFastUnmarshalMixedProto-12      | 47408    | 49790 ns/op  | 50337 B/op | 649 allocs/op  |
  | BenchmarkGoogleUnmarshalMixedProto-12    | 20776    | 115567 ns/op | 63217 B/op | 1989 allocs/op |
PASS
ok      github.com/billyplus/fastproto/test     88.420s

