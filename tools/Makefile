PROTO_DIR := proto/auth.proto
GEN_DIR := gen/auth

gen:
	protoc -I proto $(PROTO_DIR) \
	--go_out=$(GEN_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(GEN_DIR) \
	--go-grpc_opt=paths=source_relative

protoc -I proto proto/auth.proto --go_out=.gen/auth --go_opt=paths=source_relative --go-grpc_out=.gen/auth --go-grpc_opt=paths=source_relative