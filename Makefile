.PHONY: go_gen
go_gen:
	@echo "--- go generate start ---"
	@go generate ./...
	@echo "--- go generate end ---"

protoc_gen:
	@echo "--- protoc generate start ---"
	@protoc \
		--proto_path=. \
		--go_out=. \
		--go_opt=module=github.com/go-leo/cqrs \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/go-leo/cqrs \
		--go-cqrs_out=. \
		--go-cqrs_opt=module=github.com/go-leo/cqrs \
		cmd/example/api/pb/*.proto
	@echo "--- protoc generate end ---"