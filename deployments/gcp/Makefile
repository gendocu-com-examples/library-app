GCR=gcr.io/gendocu-example
GRPC_IMG=library-app-grpc
GRPCWEB_IMG=library-app-grpcweb
build:
	cd ../.. && docker build -t builder -f deployments/gcp/dockers/base.Dockerfile .
	docker build -t ${GCR}/${GRPC_IMG} -f ./dockers/grpc.Dockerfile .
	docker build -t ${GCR}/${GRPCWEB_IMG} -f ./dockers/grpcweb.Dockerfile .

deploy:
	docker push ${GCR}/${GRPC_IMG}
	gcloud run deploy --project gendocu-example --image ${GCR}/${GRPC_IMG} --platform managed --region us-central1 --max-instances 1 --memory 128Mi ${GRPC_IMG}
	docker push ${GCR}/${GRPCWEB_IMG}
	gcloud run deploy --project gendocu-example  --image ${GCR}/${GRPCWEB_IMG} --platform managed --region us-central1 --max-instances 1 --memory 128Mi  ${GRPCWEB_IMG}
