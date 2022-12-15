first:
	go build -o main1 .

second:
	go build -o main2 .

generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protobuf/mail_analysis.proto

all: first second