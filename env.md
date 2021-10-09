```shell
go mod init draw_roi

go mod tidy

go get google.golang.org/grpc

go get -a github.com/golang/protobuf/protoc-gen-go

brew install protobuf

protoc --go_out=plugins=grpc:./ ./proto/draw.proto

export SG_DRAW_FONT_PATH=./resources/font/default.ttf
export SG_APP_ADDRESS=0.0.0.0:8008
export SG_APP_PROTOCOL=tcp
```