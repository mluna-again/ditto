package main

import (
	"context"
	"ditto/cookies"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	cookies.UnimplementedCookieTravelerServer
}

var flavors = map[string]string{
	"chocolate": "vanilla",
	"vanilla":   "berries",
	"berries":   "chocolate",
}

func (s *server) SendCookie(ctx context.Context, in *cookies.Cookie) (*cookies.Cookie, error) {
	log.Printf("someone sent a %s cookie!\n", in.Flavor)

	flavor, exists := flavors[in.Flavor]
	if !exists {
		flavor = "mysterious"
	}

	return &cookies.Cookie{Flavor: flavor}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	cookies.RegisterCookieTravelerServer(grpcServer, &server{})

	grpcServer.Serve(listener)
}
