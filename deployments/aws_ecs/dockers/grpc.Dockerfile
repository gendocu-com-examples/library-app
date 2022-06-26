FROM debian:buster-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /app/grpc-server /app/grpc-server
CMD ["/app/grpc-server"]