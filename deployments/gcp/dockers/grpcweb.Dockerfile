FROM debian:buster-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /app/grpcweb-server /app/grpcweb-server
CMD ["/app/grpcweb-server"]