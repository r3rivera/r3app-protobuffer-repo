#!/bin/bash

echo "Generating the necessary code from the .proto files..."
echo "Updating GO111MODULE..."
export GO111MODULE=on
echo "PROTOC Version..."
protoc --version
echo "Generating proto codes..."

FOLDER=$1

#The argument of'plugins=grpc:.' in the --go_out tells the generator to create the necessary services
protoc -I . --go_out=plugins=grpc:. ./${FOLDER}/*.proto
echo "Complete generating the code from the .proto files..."