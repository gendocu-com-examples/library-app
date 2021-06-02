package main

import (
	sdk "git.gendocu.com/gendocu/LibraryApp.git/sdk/go"
	"github.com/gendocu-com-examples/library-app/backend/pkg"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	grpcServer := grpc.NewServer()
	srvc := pkg.NewDummyService()
	sdk.RegisterBookServiceServer(grpcServer, srvc)
	wrappedGrpc := grpcweb.WrapServer(grpcServer)
	if err := http.ListenAndServe(":8003", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		wrappedGrpc.ServeHTTP(w, req)
	})); err != nil {
		panic(err)
	}
}
