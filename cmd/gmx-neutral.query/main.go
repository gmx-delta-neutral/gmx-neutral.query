package main

import grpc_server "github.com/gmx-delta-neutral/gmx-neutral.query/internal/server/grpc"

func main() {
	grpc_server.StartServer()
}
