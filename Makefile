gen:
	PATH=$(PATH):/Users/bytedance/go/bin \
	protoc \
		--proto_path=protobuf ./protobuf/badge.proto \
		--go_out=./protobuf --go_opt=paths=source_relative \
		--go-grpc_out=./protobuf --go-grpc_opt=paths=source_relative
