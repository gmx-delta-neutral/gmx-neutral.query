package main

import (
	cfg "github.com/gmx-delta-neutral/gmx-neutral.query/internal/config"
	grpc_server "github.com/gmx-delta-neutral/gmx-neutral.query/internal/server/grpc"
)

func main() {
	config := cfg.NewConfig()
	server := grpc_server.NewServer(config)

	server.StartServer()
}
