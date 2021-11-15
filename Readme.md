# Library app

[The documentation](https://doc.gendocu.com/gendocu/api/LibraryApp). This app is a part of the [tutorial](https://blog.gendocu.com/posts/grpc-web-on-aws/). It consists of multiple parts:
- `proto` - you can find proto definition here
- `backend` - the implementation of backend in golang. It has two executables - `cmd/localserver/main.go` and `cmd/awslambda/main.go`.
- `frontend` - the simple react app, that only calls the backend
- `deployment` - the serverless framework file, that deploys the whole solution on aws lambda

# How to run

You can run environment locally using command `make run-local`. 

To deploy it to lambda you need to call `make deploy` and then replace the lambda url in the Makefile.
Then you can run react app with newly deployed backend using `run-frontend-with-aws-be`.
