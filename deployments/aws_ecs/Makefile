build:
	cd ../.. && docker build -t builder -f deployments/gcp/dockers/base.Dockerfile .
	docker build -t gendocu/library-app -f ./dockers/grpc.Dockerfile .

deploy: build
	docker push gendocu/library-app:latest
