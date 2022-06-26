package main

import (
	"log"
	"net"
	"os"

	sdk "git.gendocu.com/gendocu/LibraryApp.git/sdk/go"
	"github.com/gendocu-com-examples/library-app/backend/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("starting container")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8070"
		log.Printf("Defaulting to port %s\n", port)
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Got error: %+v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	srvc := pkg.NewDummyService()
	log.Printf("starting server at %s\n", port)
	sdk.RegisterBookServiceServer(grpcServer, srvc)
	if err := grpcServer.Serve(lis); err != nil {
		log.Println("got an error", err)
	}
}
