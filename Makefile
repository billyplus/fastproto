PROTOC=protoc
GOPATH=$(HOME)/go
ts:
	@go build -o $(GOPATH)/bin/protoc-gen-ts ./cmd/protoc-gen-ts
fastproto:
	@go build -o $(GOPATH)/bin/protoc-gen-go-fast ./cmd/protoc-gen-go-fast
testproto:
	@$(PROTOC) -I=. --go_out=./ --go_opt=module=github.com/billyplus/fastproto ./test/msg.proto ./test/outer.proto ./test/nofast.proto ./test/oneof.proto
	@$(PROTOC) -I=. --go-fast_out=./ --go-fast_opt=module=github.com/billyplus/fastproto ./test/msg.proto ./test/outer.proto ./test/oneof.proto
	@$(PROTOC) -I=. --go_out=./ --go_opt=module=github.com/billyplus/fastproto --go-fast_out=./ --go-fast_opt=module=github.com/billyplus/fastproto ./test/nomarshaler.proto
	@$(PROTOC) -I=. --go_out=./ --go_opt=module=github.com/billyplus/fastproto --go-fast_out=./ --go-fast_opt=module=github.com/billyplus/fastproto ./test/equaler.proto
bench:
	@go test -benchmem -bench=. ./test --count=1 -benchtime=1s
proto:
	@$(PROTOC) -I=. --go_out=./ --go_opt=module=github.com/billyplus/fastproto ./options/options.proto
