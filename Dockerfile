FROM alpine:latest

COPY build/kubeconfigGenerator /kubeconfigGenerator

CMD ["./kubeconfigGenerator"]
