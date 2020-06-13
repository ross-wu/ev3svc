// Package main implements a EV3 server.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ross-wu/ev3svc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/ross-wu/ev3svc/proto"
)

var (
	port = flag.Int("port", 10000, "The server port.")
)

func main() {
	flag.Parse()

	ev3, err := server.NewMockEv3Server()
	if err != nil {
		log.Fatalf("NewMockServer() error: %v", err)
	}

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEv3Server(s, ev3)
	// Register reflection for grpc_cli tool.
	reflection.Register(s)

	s.Serve(conn)
}
