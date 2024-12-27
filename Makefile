go-distributor:
	@APP_ENV==dev go run distributor/cmd/main.go
	
go-worker:
	@APP_ENV==dev go run message-broker/cmd/main.go

proto:
	@protoc ./grpc-caller/grpc/server/protofiles/*.proto --go_out=grpc-caller/grpc/server/pb --go_opt=paths=source_relative --go-grpc_out=grpc-caller/grpc/server/pb --go-grpc_opt=paths=source_relative --proto_path=grpc-caller/grpc/server/protofiles/

