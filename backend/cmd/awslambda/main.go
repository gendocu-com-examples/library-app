package main

import (
	sdk "git.gendocu.com/gendocu/LibraryApp.git/sdk/go"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	"github.com/gendocu-com-examples/library-app/backend/pkg"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	grpcServer := grpc.NewServer()
	srvc := pkg.NewDynamoDBService()
	sdk.RegisterBookServiceServer(grpcServer, srvc)
	wrappedGrpc := grpcweb.WrapServer(grpcServer)
	lambda.Start(handlerfunc.NewV2(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		wrappedGrpc.ServeHTTP(w, req)
	}).ProxyWithContext)
}
