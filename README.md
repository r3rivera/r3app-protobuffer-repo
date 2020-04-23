Protocol Buffer Storage

#Quick start
https://grpc.io/docs/quickstart/go/ 

INSTALL THE PROTO BUFFER CODE GENERATOR (PROTOC)

1. Install the PROTOC
https://github.com/protocolbuffers/protobuf/releases/tag/v3.11.4 

Install the protoc dependencies
1. Install the GO Plugin, this will be installed in $GOPATH Directory with a file protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go 

2. Add the executable protoc-gen-go to your $PATH so that protoc see can see it
export PATH="$PATH:$(go env GOPATH)/bin"

INSTALL GOLANG DEPENDENCIES
1. go get -u google.golang.org/grpc
2. go get -u github.com/golang/protobuf/protoc-gen-go

#Visual Code Plugin for protocol buffer syntax highlight
https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3


Note: Git tagging with version steps
1. git add .
2. git commit -m "message"
3. git tag -a [version] -m "[version tagg message]"
4. git push --tags

