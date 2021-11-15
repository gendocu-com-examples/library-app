.PHONY: run-local
# GRPCWEB_URL="https://t30z1m0w81.execute-api.us-east-1.amazonaws.com" # Gendocu's AWS Lambda - with DB
# GRPCWEB_URL="https://library-app-grpcweb-uqnits2f5q-uc.a.run.app" # Gendocu's GCP Cloud Run - without DB
GRPCWEB_URL="replace-me"

run-local:
	make -j run-local-backend run-local-frontend
deploy:
	go build -o deployments/bin/lambda ./backend/cmd/awslambda/main.go
	cd deployments/aws && npm install . && sls deploy
	@ echo "PLEASE COPY THE RECEIVED URL AND REPLACE THEN VARIABLE GRPCWEB_URL in Makefile"

deploy-function:
	go build -o deployments/bin/lambda ./backend/cmd/awslambda/main.go
	cd deployments/aws && sls deploy function -f hello

deploy-gcp:
	cd deployments/gcp && make build deploy

deploy-docs:
	make deploy-gcp-documentation deploy-buf-documentation
deploy-gcp-documentation:
	cd proto && gcloud endpoints services deploy api_descriptor.pb api_config.yaml --project gendocu-example
deploy-buf-documentation:
	cd proto && buf push
create-local-docs:
	cd proto && protoc --doc_out=./doc --doc_opt=html,index.html proto/*.proto

run-frontend-with-aws-be:
	@ echo "running with backend ${GRPCWEB_URL}" && sleep 3s
	cd frontend && REACT_APP_BACKEND=${GRPCWEB_URL} yarn start
run-local-backend:
	go run ./backend/cmd/localserver
run-local-frontend:
	cd frontend && npm start