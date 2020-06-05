package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kminehart/caye"
	"github.com/kminehart/caye/internal/builds"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	caye.RegisterBuildsServer(grpcServer, builds.NewServer())

	grpcServer.Serve(lis)
}
