# Experimente with grpc

### Install go module

```bash
go get -u google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Install protopuf compiler

```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v36.1/protoc-36.1-linux-x86_64.zip
unzip protoc-36.1-linux-x86_64.zip
mv protec /usr/local/bin
rm -rf include
```

### Compile .proto file

```bash
export PATH=$PATH:$(go env GOPATH)/bin
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  hello.proto
```
