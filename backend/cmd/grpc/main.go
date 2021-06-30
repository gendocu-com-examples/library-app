package main

import (
	sdk "git.gendocu.com/gendocu/LibraryApp.git/sdk/go"
	"github.com/gendocu-com-examples/library-app/backend/pkg"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	log.Println("starting container")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8070"
		log.Printf("Defaulting to port %s", port)
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Got error: %+v", err)
	}
	grpcServer := grpc.NewServer()
	srvc := pkg.NewDummyService()
	sdk.RegisterBookServiceServer(grpcServer, srvc)
	if err := grpcServer.Serve(lis); err != nil {
		log.Println("got an error", err)
	}
}
