# perl2go_grpc
Experimental project using a Golang gRPC service with HTTP gateway for prime number generation, with emphasis on Perl to 2 Go communication

## Running the Service

### 1. Start the HTTP gateway (in another terminal)

Gateway will start on port 8080
```bash
go run cmd/http_gateway/main.go
```

### 2. Start the gRPC server
Server will start on port 50051
```bash
go run cmd/grpc_server/gateway.go
```

### 3. Test the service
HTTP/REST API:
```bash
curl "http://localhost:8080/v1/primes?limit=10"
```
Expected response:
```bash
json{"primes":[2,3,5,7,11,13,17,19,23,29]}
```

### 4.Perl client:
```bash
# Install Perl dependencies (if needed)
cpan LWP::UserAgent JSON
```
Run the client
```bash
perl clientSample/clientSample.pl
```
