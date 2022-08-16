#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative auditProto.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative auditProto.proto
protoc --grpc-gateway_out=logtostderr=true:. auditProto.proto