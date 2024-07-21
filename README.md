# arb-protos
proto files for microservices usage


 protoc --proto_path=./protos --go_out=./protos --go_opt=paths=source_relative --go-grpc_out=./protos --go-grpc_opt=paths=source_relative ./protos/*/*.proto
