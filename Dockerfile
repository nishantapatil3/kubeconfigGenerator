FROM alpine:latest

COPY kubeconfigGenerator /kubeconfigGenerator

CMD ["./kubeconfigGenerator"]
