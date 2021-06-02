.PHONY: run-local
AWS_LAMBDA_URL="replace-me"

run-local:
	make -j run-local-backend run-local-frontend
deploy:
	go build -o deployments/bin/lambda ./backend/cmd/awslambda/main.go
	cd deployments && npm install . && sls deploy
	@ echo "PLEASE COPY THE RECEIVED URL AND REPLACE THEN VARIABLE AWS_LAMBDA_URL in Makefile"

deploy-function:
	go build -o deployments/bin/lambda ./backend/cmd/awslambda/main.go
	cd deployments && sls deploy function -f hello

run-frontend-with-aws-be:
	cd frontend && REACT_APP_BACKEND=${AWS_LAMBDA_URL} yarn start
run-local-backend:
	go run ./backend/cmd/localserver
run-local-frontend:
	cd frontend && npm start