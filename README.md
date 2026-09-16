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

genereierter code nach **gen/go** Proto file nach **proto/hello/v1**

```bash
export PATH=$PATH:$(go env GOPATH)/bin
protoc \
  --go_out=gen/go \
  --go_opt=paths=source_relative \
  --go-grpc_out=gen/go \
  --go-grpc_opt=paths=source_relative \
  proto/hello/v1/hello.proto
```

### SSL

```bash
openssl req \
  -x509 \
  -newkey rsa:4096 \
  -keyout server.key \
  -out server.crt \
  -days 365 \
  -nodes \
  -subj "/CN=localhost"
```

## MTLS

### CA

```bash
openssl genrsa -out ca.key 4096

openssl req -x509 \
  -new \
  -nodes \
  -key ca.key \
  -sha256 \
  -days 3650 \
  -out ca.crt \
  -subj "/CN=Test CA"
```

### Server

```bash
openssl genrsa -out server.key 4096

openssl req -new \
  -key server.key \
  -out server.csr \
  -subj "/CN=grpc-server"

openssl x509 -req \
  -in server.csr \
  -CA ca.crt \
  -CAkey ca.key \
  -CAcreateserial \
  -out server.crt \
  -days 365
```

### Client

```bash
openssl genrsa -out client.key 4096

openssl req -new \
  -key client.key \
  -out client.csr \
  -subj "/CN=grpc-client"

openssl x509 -req \
  -in client.csr \
  -CA ca.crt \
  -CAkey ca.key \
  -CAcreateserial \
  -out client.crt \
  -days 365
```

## MTLS Kubernetes cert-namanager

### Issuer

```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: selfsigned-root
spec:
  selfSigned: {}
```

### Root CA

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: grpc-root-ca
  namespace: grpctest
spec:
  isCA: true
  commonName: grpc-root-ca
  secretName: grpc-root-ca-secret
  duration: 87600h # 10 Jahre
  privateKey:
    algorithm: RSA
    size: 4096
  issuerRef:
    name: selfsigned-root
    kind: ClusterIssuer
```

### CA Issuer

```yaml
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: grpc-ca-issuer
  namespace: grpctest
spec:
  ca:
    secretName: grpc-root-ca-secret
```

### Server Zertifikat

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: grpc-server
  namespace: grpctest
spec:
  secretName: grpc-server-tls

  commonName: grpc-server.grpctest.svc.cluster.local

  dnsNames:
    - grpc-server
    - grpc-server.grpctest
    - grpc-server.grpctest.svc
    - grpc-server.grpctest.svc.cluster.local

  privateKey:
    algorithm: RSA
    size: 4096

  issuerRef:
    name: grpc-ca-issuer
    kind: Issuer
```

### Client

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: grpc-client
  namespace: grpctest
spec:
  secretName: grpc-client-tls

  commonName: grpc-client

  usages:
    - client auth

  privateKey:
    algorithm: RSA
    size: 4096

  issuerRef:
    name: grpc-ca-issuer
    kind: Issuer
```
