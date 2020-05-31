# proto generate
* 1. cd ../github.com
* 2. protoc --proto_path=. --go_out=plugins=grpc:. github.com/wangrenyi/kmpgo/proto/*.proto

# example
* protoc --proto_path=. --go_out=plugins=grpc:. github.com/wangrenyi/kmpgo/proto/common/common.proto