GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/server -v cmd/zmq/server/zmq_server.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/client -v cmd/zmq/client/zmq_client.go
