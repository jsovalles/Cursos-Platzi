package main

import (
	"github.com/jsovalles/grpc/repository"
	"github.com/jsovalles/grpc/server"
	"github.com/jsovalles/grpc/testpb"
	"github.com/jsovalles/grpc/utils"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5070")

	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}

	env := utils.NewEnv()
	database, err := utils.NewDatabase(env)
	if err != nil {
		log.Fatalf("Error with database: %s", err.Error())
	}

	repo := repository.NewRepository(env, database)

	server := server.NewTestServer(repo)

	if err != nil {
		log.Fatalf("Error creating repository: %s", err.Error())
	}

	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatalf("Error serving: %s", err.Error())
	}
}
