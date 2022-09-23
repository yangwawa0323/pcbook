protoc_gen:
	protoc --proto_path=proto proto/*.proto --go_out=./pb \
	--go_opt=paths=source_relative   --go-grpc_out=./pb \
	--go-grpc_opt=paths=source_relative

gorm:
	protoc --proto_path=proto \
	--proto_path=third_party/proto  \
	proto/*.proto --go_out=./pb \
	--go_opt=paths=source_relative --go-grpc_out=./pb \
	--go-grpc_opt=paths=source_relative \
	--gorm_out="engine=mysql:./pb"

	--gorm_opt=paths=source_relative

buf_gen:
	buf generate

inject:
	protoc-go-inject-tag -input=./pb/*.pb.go

clean:
	rm ./pb/*.pb.go
	rm ./pb/*.gorm.go
	rm ./pb/* -rf

run:
	go run main.go

test:
	go test -cover ./...

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

gui:
	grpcui -plaintext localhost:8080
