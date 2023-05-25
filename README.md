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

### Benchmarks

[code](https://github.com/billyplus/fastproto/tree/main/test)

```
goos: linux
goarch: amd64
pkg: github.com/billyplus/fastproto/test
cpu: AMD Ryzen 9 5950X 16-Core Processor

BenchmarkFastMarshalStringSlice-6         	14644707	        81.75 ns/op	      80 B/op	       1 allocs/op
BenchmarkGoogleMarshalStringSlice-6       	 8223910	       144.4 ns/op	      80 B/op	       1 allocs/op

BenchmarkFastMarshalBytesSlice-6          	13065022	        93.40 ns/op	      80 B/op	       1 allocs/op
BenchmarkGoogleMarshalBytesSlice-6        	10043254	       124.9 ns/op	      80 B/op	       1 allocs/op

BenchmarkFastMarshalInt32Slice-6          	 5772819	       213.1 ns/op	     128 B/op	       1 allocs/op
BenchmarkGoogleMarshalInt32Slice-6        	 5056791	       237.5 ns/op	     128 B/op	       1 allocs/op

BenchmarkFastMarshalSint64Slice-6         	 4123633	       288.3 ns/op	     224 B/op	       1 allocs/op
BenchmarkGoogleMarshalSint64Slice-6       	 3811389	       311.4 ns/op	     224 B/op	       1 allocs/op

BenchmarkFastMarshalSfixed32Slice-6       	16257074	        73.97 ns/op	     112 B/op	       1 allocs/op
BenchmarkGoogleMarshalSfixed32Slice-6     	12917850	        93.63 ns/op	     112 B/op	       1 allocs/op

BenchmarkFastMarshalSfixed64Slice-6       	14003510	        89.69 ns/op	     208 B/op	       1 allocs/op
BenchmarkGoogleMarshalSfixed64Slice-6     	11058189	       115.9 ns/op	     208 B/op	       1 allocs/op

BenchmarkFastMarshalToMixedProto-6        	   74734	     15354 ns/op	       0 B/op	       0 allocs/op
BenchmarkFastMarshalMixedProto-6          	   43844	     27804 ns/op	   18432 B/op	       1 allocs/op
BenchmarkGoogleMarshalMixedProto-6        	   12552	     94428 ns/op	   37664 B/op	    1521 allocs/op

BenchmarkFastSizeMixedProto-6             	  205432	      6061 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoogleSizeMixedProto-6           	   32412	     39230 ns/op	    9616 B/op	     760 allocs/op

BenchmarkFastUnmarshalStringSlice-6       	 4322337	       291.3 ns/op	     314 B/op	       7 allocs/op
BenchmarkGoogleUnmarshalStringSlice-6     	 3088686	       384.5 ns/op	     314 B/op	       7 allocs/op

BenchmarkFastUnmarshalBytesSlice-6        	 3194150	       376.0 ns/op	     448 B/op	       8 allocs/op
BenchmarkGoogleUnmarshalBytesSlice-6      	 2770154	       426.6 ns/op	     448 B/op	       8 allocs/op

BenchmarkFastUnmarshalInt32Slice-6        	 6377149	       183.2 ns/op	     112 B/op	       1 allocs/op
BenchmarkGoogleUnmarshalInt32Slice-6      	 3752682	       318.7 ns/op	     248 B/op	       5 allocs/op

BenchmarkFastUnmarshalSint64Slice-6       	 4416526	       271.5 ns/op	     208 B/op	       1 allocs/op
BenchmarkGoogleUnmarshalSint64Slice-6     	 2903524	       405.0 ns/op	     504 B/op	       6 allocs/op

BenchmarkFastUnmarshalSfixed32Slice-6     	14313001	        85.07 ns/op	     112 B/op	       1 allocs/op
BenchmarkGoogleUnmarshalSfixed32Slice-6   	 5353230	       224.2 ns/op	     248 B/op	       5 allocs/op

BenchmarkFastUnmarshalSfixed64Slice-6     	12808696	       103.0 ns/op	     208 B/op	       1 allocs/op
BenchmarkGoogleUnmarshalSfixed64Slice-6   	 3824290	       317.3 ns/op	     504 B/op	       6 allocs/op

BenchmarkFastUnmarshalMixedProto-6        	   20580	     58110 ns/op	   46909 B/op	     606 allocs/op
BenchmarkGoogleUnmarshalMixedProto-6      	    8949	    132525 ns/op	   60842 B/op	    1966 allocs/op
```


